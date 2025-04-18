package service

import (
	"github.com/teamdetected/internal/model"
	"github.com/teamdetected/internal/repository"
)

type Service struct {
	Authorization
}

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(email, password string) (model.User, error)
	GenerateToken(email, password string) (string, error)
	DeleteUser(id int) error
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
