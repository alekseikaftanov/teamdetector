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
	GenerateToken(email, password string) (string, error)
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
