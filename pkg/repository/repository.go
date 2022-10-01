package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Admin interface {
}

type Repository struct {
	Authorization
	Admin
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
