package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Video VideoModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Video: VideoModel{DB: db},
	}
}
