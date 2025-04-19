package service

import (
	"fmt"
	"log"
	"net/smtp"
)

type EmailService struct {
	from     string
	password string
	host     string
	port     string
}

func NewEmailService(from, password, host, port string) *EmailService {
	return &EmailService{
		from:     from,
		password: password,
		host:     host,
		port:     port,
	}
}

func (s *EmailService) SendSurveyInvitation(email, name string, teamID int) error {
	// Формируем ссылку на опрос
	surveyLink := fmt.Sprintf("http://localhost:8080/surveys/team/%d", teamID)

	// Формируем текст письма
	subject := "Приглашение пройти опрос"
	body := fmt.Sprintf(`Здравствуйте, %s!

Вы были добавлены в команду. Пожалуйста, пройдите опрос по ссылке:
%s

С уважением,
Команда TeamDetected`, name, surveyLink)

	// Формируем MIME сообщение
	msg := fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/plain; charset=\"UTF-8\"\r\n\r\n"+
		"%s", s.from, email, subject, body)

	// Настраиваем аутентификацию
	auth := smtp.PlainAuth("", s.from, s.password, s.host)
	addr := fmt.Sprintf("%s:%s", s.host, s.port)

	log.Printf("Attempting to send email to %s via %s", email, addr)

	// Отправляем email
	err := smtp.SendMail(addr, auth, s.from, []string{email}, []byte(msg))
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return fmt.Errorf("failed to send survey invitation: %v", err)
	}

	log.Printf("Email sent successfully to %s", email)
	return nil
}
