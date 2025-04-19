package service

import (
	"github.com/teamdetected/internal/model"
	"github.com/teamdetected/internal/repository"
)

type Service struct {
	Authorization
	Company
	Team
	Survey
}

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(email, password string) (model.User, error)
	GenerateToken(email, password string) (string, error)
	DeleteUser(id int) error
}

type Company interface {
	CreateCompany(company model.Company) (int, error)
	GetCompanyByID(id int) (model.Company, error)
	GetCompaniesByUserID(userID int) ([]model.Company, error)
	DeleteCompany(id int) error
}

type Team interface {
	CreateTeam(team model.Team) (int, error)
	GetTeamByID(id int) (model.Team, error)
	GetTeamsByCompanyID(companyID int) ([]model.Team, error)
	DeleteTeam(id int) error
}

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

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Company:       NewCompanyService(repos.Company),
		Team:          NewTeamService(repos.Team),
		Survey:        NewSurveyService(repos.Survey),
	}
}
