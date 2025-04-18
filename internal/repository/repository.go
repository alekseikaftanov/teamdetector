package repository

import (
	"database/sql"

	"github.com/teamdetected/internal/model"
)

type Repository struct {
	Authorization
}

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(email, password string) (model.User, error)
	DeleteUser(id int) error
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
