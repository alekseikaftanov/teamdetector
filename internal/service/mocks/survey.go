package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/teamdetected/internal/model"
)

type Survey struct {
	mock.Mock
}

func (m *Survey) CreateSurvey(survey model.Survey) (int, error) {
	args := m.Called(survey)
	return args.Int(0), args.Error(1)
}

func (m *Survey) GetSurveyByID(id int) (model.Survey, error) {
	args := m.Called(id)
	return args.Get(0).(model.Survey), args.Error(1)
}

func (m *Survey) GetSurveysByTeamID(teamID int) ([]model.Survey, error) {
	args := m.Called(teamID)
	return args.Get(0).([]model.Survey), args.Error(1)
}

func (m *Survey) DeleteSurvey(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *Survey) CreateSurveyResponse(response model.SurveyResponse) (int, error) {
	args := m.Called(response)
	return args.Int(0), args.Error(1)
}

func (m *Survey) GetSurveyResponses(surveyID int) ([]model.SurveyResponse, error) {
	args := m.Called(surveyID)
	return args.Get(0).([]model.SurveyResponse), args.Error(1)
}

func (m *Survey) GetSurveyOptions() ([]model.SurveyOption, error) {
	args := m.Called()
	return args.Get(0).([]model.SurveyOption), args.Error(1)
}

func (m *Survey) GetSurveyQuestions() ([]model.SurveyQuestion, error) {
	args := m.Called()
	return args.Get(0).([]model.SurveyQuestion), args.Error(1)
}
