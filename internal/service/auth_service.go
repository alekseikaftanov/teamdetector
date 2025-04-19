package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/teamdetected/internal/model"
	"github.com/teamdetected/internal/repository"
	"golang.org/x/crypto/bcrypt"
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	user.Password = string(hashedPassword)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GetUser(email, password string) (model.User, error) {
	return s.repo.GetUser(email, password)
}

func (s *AuthService) GetUserByID(id int) (model.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *AuthService) UpdateUser(id int, input model.UpdateUserInput) error {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return err
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Role = input.Role
	user.UpdatedAt = time.Now()

	return s.repo.UpdateUser(id, user)
}

func (s *AuthService) ChangePassword(id int, oldPassword, newPassword string) error {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return err
	}

	// Проверяем старый пароль
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return errors.New("incorrect old password")
	}

	// Хешируем новый пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	user.UpdatedAt = time.Now()

	return s.repo.UpdateUser(id, user)
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

func (s *AuthService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}
