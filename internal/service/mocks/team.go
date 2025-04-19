package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/teamdetected/internal/model"
)

type Team struct {
	mock.Mock
}

func NewTeam(t mock.TestingT) *Team {
	return &Team{}
}

func (m *Team) CreateTeam(team model.Team) (int, error) {
	args := m.Called(team)
	return args.Int(0), args.Error(1)
}

func (m *Team) GetTeamByID(id int) (model.Team, error) {
	args := m.Called(id)
	return args.Get(0).(model.Team), args.Error(1)
}

func (m *Team) GetTeamsByCompanyID(companyID int) ([]model.Team, error) {
	args := m.Called(companyID)
	return args.Get(0).([]model.Team), args.Error(1)
}

func (m *Team) DeleteTeam(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
