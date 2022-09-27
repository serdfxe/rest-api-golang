package repository

type Authorization interface {
}

type Admin interface {
}

type Repository struct {
	Authorization
	Admin
}

func NewRepository() *Repository {
	return &Repository{}
}
