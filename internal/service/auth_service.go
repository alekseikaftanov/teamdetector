package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/teamdetected/internal/model"
	"github.com/teamdetected/internal/repository"
)

const (
	tokenTTL   = 12 * time.Hour
	signingKey = "your-secret-key" // В продакшене нужно использовать переменную окружения
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) (int, error) {
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.repo.GetUser(email, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(tokenTTL).Unix(),
	})

	return token.SignedString([]byte(signingKey))
}
