package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/strimenko/auth"
)

type Authorization interface {
	CreateUser(user auth.User) (int, error)
	GetUser(username, password string) (auth.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
