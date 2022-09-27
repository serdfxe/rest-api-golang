package service

import "github.com/serdfxe/rest-api-golang/pkg/repository"

type Authorization interface {
}

type Admin interface {
}

type Service struct {
	Authorization
	Admin
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
