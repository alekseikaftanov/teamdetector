package service

import (
	"github.com/teamdetected/internal/model"
	"github.com/teamdetected/internal/repository"
)

type SurveyService struct {
	repo repository.Survey
}

func NewSurveyService(repo repository.Survey) *SurveyService {
	return &SurveyService{repo: repo}
}

func (s *SurveyService) CreateSurvey(survey model.Survey) (int, error) {
	if survey.TeamID == 0 {
		return 0, model.ErrInvalidInput
	}
	if survey.CreatedBy == 0 {
		return 0, model.ErrInvalidInput
	}

	survey.Status = "active"
	return s.repo.CreateSurvey(survey)
}

func (s *SurveyService) GetSurveyByID(id int) (model.Survey, error) {
	return s.repo.GetSurveyByID(id)
}

func (s *SurveyService) GetSurveysByTeamID(teamID int) ([]model.Survey, error) {
	return s.repo.GetSurveysByTeamID(teamID)
}

func (s *SurveyService) DeleteSurvey(id int) error {
	return s.repo.DeleteSurvey(id)
}

func (s *SurveyService) CreateSurveyResponse(response model.SurveyResponse) (int, error) {
	if response.SurveyID == 0 || response.UserID == 0 || response.QuestionID == 0 || response.OptionID == 0 {
		return 0, model.ErrInvalidInput
	}
	return s.repo.CreateSurveyResponse(response)
}

func (s *SurveyService) GetSurveyResponses(surveyID int) ([]model.SurveyResponse, error) {
	return s.repo.GetSurveyResponses(surveyID)
}

func (s *SurveyService) GetSurveyOptions() ([]model.SurveyOption, error) {
	return s.repo.GetSurveyOptions()
}

func (s *SurveyService) GetSurveyQuestions() ([]model.SurveyQuestion, error) {
	return s.repo.GetSurveyQuestions()
}
