package service

import (
	rest "github.com/serdfxe/rest-api-golang/pkg"
	"github.com/serdfxe/rest-api-golang/pkg/repository"
)

type Authorization interface {
	CreateUser(user rest.User) (int, error)
}

type Admin interface {
}

type Service struct {
	Authorization
	Admin
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repos.Authorization),
	}
}
