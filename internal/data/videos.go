package data

import (
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

func (v VideoModel) update(video *Video) error {
	return nil
}

func (v VideoModel) Delete(id int64) error {
	return nil
}
