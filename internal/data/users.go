package data

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var ErrDuplicateEmail = errors.New("duplicate email")

var AnonymousUser = &User{}

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  password  `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) IsAnonymous() bool {
	return u == AnonymousUser
}

type password struct {
	plaintext *string
	hash      []byte
}

func (p *password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}
	p.plaintext = &plaintextPassword
	p.hash = hash
	return nil
}

func (p *password) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}

type UserModel struct {
	DB *sql.DB
}

func (m UserModel) Insert(user *User) error {
	query := `INSERT INTO users (name, email, password_hash, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, created_at, updated_at`
	args := []any{user.Name, user.Email, user.Password.hash, user.CreatedAt, user.UpdatedAt}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}

	return nil
}

func (m UserModel) GetByEmail(email string) (*User, error) {
	query := `SELECT id, name, email, password_hash, created_at, updated_at
	FROM users
	WHERE email = $1`
	var user User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password.hash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &user, nil
}

func (m UserModel) Update(user *User) error {
	query := `UPDATE users
	SET name = $1, updated_at = $2
	WHERE email = $3`
	args := []any{user.Name, user.UpdatedAt, user.Email}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rs, err := m.DB.ExecContext(ctx, query, args...)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := rs.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rowsAffected == 0 {
		return ErrEditConflict
	}

	return nil
}

func (m UserModel) GetByToken(tokenPlaintext string) (*User, error) {
	tokenHash := sha256.Sum256([]byte(tokenPlaintext))
	query := `SELECT users.id, users.name, users.email, users.password_hash, users.created_at, users.updated_at
	FROM users
	JOIN tokens
	ON users.id = tokens.user_id
	WHERE tokens.hash = $1`
	args := []any{tokenHash[:]}
	var user User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.Name, &user.Email, &user.Password.hash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &user, nil
}
