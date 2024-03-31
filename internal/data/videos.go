package data

import "time"

type Video struct {
	ID    int64 `json:"id"`
	Url   string `json:"url"`
	Title string `json:"title"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}