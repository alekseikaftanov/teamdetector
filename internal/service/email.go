package service

import (
	"fmt"
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
	surveyLink := fmt.Sprintf("http://your-domain.com/surveys/team/%d", teamID)

	// Формируем текст письма
	subject := "Приглашение пройти опрос"
	body := fmt.Sprintf(`Здравствуйте, %s!

Вы были добавлены в команду. Пожалуйста, пройдите опрос по ссылке:
%s

С уважением,
Команда TeamDetected`, name, surveyLink)

	msg := fmt.Sprintf("Subject: %s\n\n%s", subject, body)

	// Отправляем email
	auth := smtp.PlainAuth("", s.from, s.password, s.host)
	addr := fmt.Sprintf("%s:%s", s.host, s.port)

	return smtp.SendMail(addr, auth, s.from, []string{email}, []byte(msg))
}
