package data

import (
	"testing"

	"youshare-api.anvo.dev/internal/assert"
)

func TestUserModel_GetByEmail(t *testing.T) {
	testCases := []struct {
		name     string
		email    string
		expected *User
	}{
		{
			name:  "valid email",
			email: "alice_test@example.com",
			expected: &User{
				Name:  "Alice",
				Email: "alice_test@example.com",
			},
		},
		{
			name:  "invalid email",
			email: "invalid@example",
			expected: &User{
				Name:  "",
				Email: "",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db := newTestDB(t)

			m := UserModel{db}

			user, err := m.GetByEmail(tc.email)

			if tc.name == "valid email" {
				assert.Equal(t, tc.expected.Email, user.Email)
				assert.Equal(t, tc.expected.Name, user.Name)
			}

			if tc.name == "invalid email" {
				assert.Equal(t, ErrRecordNotFound, err)
			}
		})
	}
}

func TestUserModel_Insert(t *testing.T) {
	testCases := []struct {
		name          string
		user          *User
		expectedError error
	}{
		{
			name: "valid user",
			user: &User{
				Name:  "Alice",
				Email: "foo@example.com",
			},
			expectedError: nil,
		},
		{
			name: "duplicate email",
			user: &User{
				Name:  "Alice",
				Email: "alice_test@example.com", // same email as seed test user
			},
			expectedError: ErrDuplicateEmail,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db := newTestDB(t)
			m := UserModel{db}

			err := m.Insert(tc.user)
			assert.Equal(t, tc.expectedError, err)

			if tc.expectedError == nil {
				var user User
				err = db.QueryRow("SELECT id, name, email, created_at, updated_at FROM users WHERE email = $1", tc.user.Email).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
				assert.NilError(t, err)
				assert.Equal(t, tc.user.Name, user.Name)
				assert.Equal(t, tc.user.Email, user.Email)
			}
		})
	}
}

func TestUserModel_Update(t *testing.T) {
	testCases := []struct {
		ID            int64
		name          string
		user          *User
		expectedError error
		expectedUser  *User
	}{
		{
			name: "valid user",
			user: &User{
				Name:  "Alice updated",
				Email: "alice_test@example.com",
			},
			expectedError: nil,
		},
		{
			name: "nonexistent email",
			user: &User{
				Name:  "Alice updated",
				Email: "invalid@example.com",
			},
			expectedError: ErrEditConflict,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db := newTestDB(t)
			m := UserModel{db}

			err := m.Update(tc.user)
			assert.Equal(t, tc.expectedError, err)

			if tc.expectedError == nil {
				var user User
				err = db.QueryRow("SELECT id, name, email, created_at, updated_at FROM users WHERE email = $1", tc.user.Email).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
				assert.NilError(t, err)
				assert.Equal(t, tc.user.Name, user.Name)
				assert.Equal(t, tc.user.Email, user.Email)
			}
		})
	}
}
