package repository

import (
	"github.com/jmoiron/sqlx"
	rest "github.com/serdfxe/rest-api-golang/pkg"
)

type Authorization interface {
	CreateUser(user rest.User) (int, error)
	GetUser(username string, password string) (rest.User, error)
}

type Admin interface {
}

type Repository struct {
	Authorization
	Admin
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: newAuthPostgres(db),
	}
}
