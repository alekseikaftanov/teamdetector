package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/teamdetected/internal/model"
)

type Company struct {
	mock.Mock
}

func NewCompany(t mock.TestingT) *Company {
	return &Company{}
}

func (m *Company) CreateCompany(company model.Company) (int, error) {
	args := m.Called(company)
	return args.Int(0), args.Error(1)
}

func (m *Company) GetCompanyByID(id int) (model.Company, error) {
	args := m.Called(id)
	return args.Get(0).(model.Company), args.Error(1)
}

func (m *Company) GetCompaniesByUserID(userID int) ([]model.Company, error) {
	args := m.Called(userID)
	return args.Get(0).([]model.Company), args.Error(1)
}

func (m *Company) DeleteCompany(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
