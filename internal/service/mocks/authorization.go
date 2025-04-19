package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/teamdetected/internal/model"
)

type Authorization struct {
	mock.Mock
}

func NewAuthorization(t mock.TestingT) *Authorization {
	return &Authorization{}
}

func (m *Authorization) CreateUser(user model.User) (int, error) {
	args := m.Called(user)
	return args.Int(0), args.Error(1)
}

func (m *Authorization) GetUser(email, password string) (model.User, error) {
	args := m.Called(email, password)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *Authorization) GenerateToken(email, password string) (string, error) {
	args := m.Called(email, password)
	return args.String(0), args.Error(1)
}

func (m *Authorization) DeleteUser(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
