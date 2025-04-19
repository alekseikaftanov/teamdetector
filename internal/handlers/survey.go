package handlers

import (
	"github.com/jmoiron/sqlx"
	"github.com/teamdetected/internal/repository"
	"github.com/teamdetected/internal/service"
)

type SurveyHandler struct {
	service *service.SurveyService
	db      *sqlx.DB
}

func NewSurveyHandler(db *sqlx.DB) *SurveyHandler {
	repo := repository.NewSurveyPostgres(db)
	return &SurveyHandler{
		service: service.NewSurveyService(repo),
		db:      db,
	}
}

// ... existing code ...
