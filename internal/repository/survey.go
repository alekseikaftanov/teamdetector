package repository

import "github.com/teamdetected/internal/model"

type Survey interface {
	CreateSurvey(survey model.Survey) (int, error)
	GetSurveyByID(id int) (model.Survey, error)
	GetSurveysByTeamID(teamID int) ([]model.Survey, error)
	DeleteSurvey(id int) error
	CreateSurveyResponse(response model.SurveyResponse) (int, error)
	GetSurveyResponses(surveyID int) ([]model.SurveyResponse, error)
	GetSurveyOptions() ([]model.SurveyOption, error)
	GetSurveyQuestions() ([]model.SurveyQuestion, error)
}
