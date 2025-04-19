package service

import (
	"database/sql"
	"errors"
	"log"
	"math/rand"

	"github.com/teamdetected/internal/model"
	"github.com/teamdetected/internal/repository"
)

type TeamService struct {
	repo     repository.Team
	authRepo repository.Authorization
	email    Email
}

func NewTeamService(repo repository.Team, authRepo repository.Authorization, email Email) *TeamService {
	return &TeamService{
		repo:     repo,
		authRepo: authRepo,
		email:    email,
	}
}

func (s *TeamService) CreateTeam(team model.Team) (int, error) {
	return s.repo.CreateTeam(team)
}

func (s *TeamService) GetTeamByID(id int) (model.Team, error) {
	return s.repo.GetTeamByID(id)
}

func (s *TeamService) GetTeamsByCompanyID(companyID int) ([]model.Team, error) {
	return s.repo.GetTeamsByCompanyID(companyID)
}

func (s *TeamService) DeleteTeam(id int) error {
	return s.repo.DeleteTeam(id)
}

func (s *TeamService) AddUserToTeam(teamID int, input model.AddUserToTeamInput) error {
	// Проверяем существование команды
	if _, err := s.repo.GetTeamByID(teamID); err != nil {
		return err
	}

	// Получаем или создаем пользователя
	user, err := s.authRepo.GetUserByEmail(input.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Создаем временный пароль
			tempPassword := generateTempPassword()

			// Создаем нового пользователя
			user = model.User{
				Email:    input.Email,
				Name:     input.Name,
				Role:     "member",
				Password: tempPassword,
			}
			userID, err := s.authRepo.CreateUser(user)
			if err != nil {
				return err
			}
			user.ID = userID

			// Отправляем приглашение с временным паролем
			if err := s.email.SendSurveyInvitation(user.Email, user.Name, teamID); err != nil {
				// Логируем ошибку, но не прерываем выполнение
				log.Printf("Failed to send survey invitation: %v", err)
			}
		} else {
			return err
		}
	}

	// Добавляем пользователя в команду
	return s.repo.AddUserToTeam(teamID, user.ID)
}

func (s *TeamService) AddUsersToTeam(teamID int, inputs []model.AddUserToTeamInput) error {
	// Проверяем существование команды
	if _, err := s.repo.GetTeamByID(teamID); err != nil {
		return err
	}

	// Начинаем транзакцию
	tx, err := s.repo.BeginTx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, input := range inputs {
		// Получаем или создаем пользователя
		user, err := s.authRepo.GetUserByEmail(input.Email)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				// Создаем временный пароль
				tempPassword := generateTempPassword()

				// Создаем нового пользователя
				user = model.User{
					Email:    input.Email,
					Name:     input.Name,
					Role:     "member",
					Password: tempPassword,
				}
				userID, err := s.authRepo.CreateUser(user)
				if err != nil {
					return err
				}
				user.ID = userID

				// Отправляем приглашение с временным паролем
				if err := s.email.SendSurveyInvitation(user.Email, user.Name, teamID); err != nil {
					// Логируем ошибку, но не прерываем выполнение
					log.Printf("Failed to send survey invitation: %v", err)
				}
			} else {
				return err
			}
		}

		// Добавляем пользователя в команду
		if err := s.repo.AddUserToTeam(teamID, user.ID); err != nil {
			return err
		}
	}

	// Подтверждаем транзакцию
	return tx.Commit()
}

// Вспомогательная функция для генерации временного пароля
func generateTempPassword() string {
	const length = 12
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
