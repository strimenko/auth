package service

import (
	"github.com/strimenko/auth"
	"github.com/strimenko/auth/pkg/repository"
)

type Authorization interface {
	CreateUser(user auth.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
