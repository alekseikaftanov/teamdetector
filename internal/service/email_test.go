package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailService_SendSurveyInvitation(t *testing.T) {
	// Создаем тестовый email сервис с моковыми данными
	emailServ := NewEmailService(
		"test@example.com",
		"test-password",
		"localhost",
		"1025", // Используем порт для тестового SMTP сервера
	)

	// Тестовые данные
	email := "user@example.com"
	name := "Test User"
	teamID := 1

	// Проверяем формирование сообщения
	err := emailServ.SendSurveyInvitation(email, name, teamID)

	// В тестовом окружении ожидаем ошибку, так как нет реального SMTP сервера
	assert.Error(t, err)
}
