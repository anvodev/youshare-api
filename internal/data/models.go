package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Videos VideoModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Videos: VideoModel{DB: db},
	}
}
