package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Video struct {
	ID          int64     `json:"id"`
	Url         string    `json:"url"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type VideoModel struct {
	DB *sql.DB
}

func (v VideoModel) Insert(video *Video) error {
	query := `
		INSERT INTO videos (url, title, description)
		VALUES ($1, $2, $3)
		RETURNING id, created_at`
	args := []any{video.Title, video.Url, video.Description}
	return v.DB.QueryRow(query, args...).Scan(&video.ID, &video.CreatedAt)
}

func (v VideoModel) GetAll() ([]*Video, error) {
	query := `
		SELECT id, url, title, description, created_at, updated_at
		FROM videos
		ORDER BY id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := v.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	videos := []*Video{}
	for rows.Next() {
		var video Video
		err := rows.Scan(&video.ID, &video.Url, &video.Title, &video.Description, &video.CreatedAt, &video.UpdatedAt)
		if err != nil {
			return nil, err
		}
		videos = append(videos, &video)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return videos, nil
}

func (v VideoModel) Get(id int64) (*Video, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
		SELECT id, url, title, description, created_at, updated_at
		FROM videos
		WHERE id = $1`
	var video Video

	err := v.DB.QueryRow(query, id).Scan(&video.ID, &video.Url, &video.Title, &video.Description, &video.CreatedAt, &video.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &video, nil
}

func (v VideoModel) Update(video *Video) error {
	query := `
		UPDATE videos
		SET title = $1, url = $2, description = $3, updated_at = $4
		WHERE id = $5`

	args := []any{video.Title, video.Url, video.Description, time.Now(), video.ID}
	_, err := v.DB.Exec(query, args...)

	return err
}

func (v VideoModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `DELETE FROM videos WHERE id = $1`
	result, err := v.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
