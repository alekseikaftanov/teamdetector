package handler

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/teamdetected/internal/model"
	"github.com/teamdetected/internal/service"
	"github.com/teamdetected/internal/service/mocks"
)

func TestHandler_Register(t *testing.T) {
	type mockBehavior func(s *mocks.Authorization, user model.User)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           model.User
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name": "Test", "email": "test@test.com", "password": "qwerty", "role": "user"}`,
			inputUser: model.User{
				Name:     "Test",
				Email:    "test@test.com",
				Password: "qwerty",
				Role:     "user",
			},
			mockBehavior: func(s *mocks.Authorization, user model.User) {
				s.On("CreateUser", user).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:      "Empty Fields",
			inputBody: `{}`,
			mockBehavior: func(s *mocks.Authorization, user model.User) {
			},
			expectedStatusCode:  400,
			expectedRequestBody: `{"error":"Key: 'SignUpInput.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'SignUpInput.Email' Error:Field validation for 'Email' failed on the 'required' tag\nKey: 'SignUpInput.Password' Error:Field validation for 'Password' failed on the 'required' tag\nKey: 'SignUpInput.Role' Error:Field validation for 'Role' failed on the 'required' tag"}`,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := new(mocks.Authorization)
			test.mockBehavior(c, test.inputUser)

			services := &service.Service{Authorization: c}
			handler := NewHandler(services)

			// Init Endpoint
			r := gin.New()
			r.POST("/api/v1/auth/register", handler.Register)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/auth/register",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_Login(t *testing.T) {
	type mockBehavior func(s *mocks.Authorization, email, password string)

	testTable := []struct {
		name                string
		inputBody           string
		inputEmail          string
		inputPassword       string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:          "OK",
			inputBody:     `{"email": "test@test.com", "password": "qwerty"}`,
			inputEmail:    "test@test.com",
			inputPassword: "qwerty",
			mockBehavior: func(s *mocks.Authorization, email, password string) {
				s.On("GenerateToken", email, password).Return("token", nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"token":"token"}`,
		},
		{
			name:      "Wrong Input",
			inputBody: `{"email": "test"}`,
			mockBehavior: func(s *mocks.Authorization, email, password string) {
			},
			expectedStatusCode:  400,
			expectedRequestBody: `{"error":"Key: 'SignInInput.Email' Error:Field validation for 'Email' failed on the 'email' tag\nKey: 'SignInInput.Password' Error:Field validation for 'Password' failed on the 'required' tag"}`,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := new(mocks.Authorization)
			test.mockBehavior(c, test.inputEmail, test.inputPassword)

			services := &service.Service{Authorization: c}
			handler := NewHandler(services)

			// Init Endpoint
			r := gin.New()
			r.POST("/api/v1/auth/login", handler.Login)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/auth/login",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_DeleteUser(t *testing.T) {
	type mockBehavior func(s *mocks.Authorization, id int)

	testTable := []struct {
		name                string
		userID              string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:   "OK",
			userID: "1",
			mockBehavior: func(s *mocks.Authorization, id int) {
				s.On("DeleteUser", id).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"message":"user deleted successfully"}`,
		},
		{
			name:   "Invalid ID",
			userID: "invalid",
			mockBehavior: func(s *mocks.Authorization, id int) {
			},
			expectedStatusCode:  400,
			expectedRequestBody: `{"error":"invalid id"}`,
		},
		{
			name:   "Service Error",
			userID: "1",
			mockBehavior: func(s *mocks.Authorization, id int) {
				s.On("DeleteUser", id).Return(errors.New("service error"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"error":"service error"}`,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := new(mocks.Authorization)
			test.mockBehavior(c, 1)

			services := &service.Service{Authorization: c}
			handler := NewHandler(services)

			// Init Endpoint
			r := gin.New()
			r.DELETE("/api/v1/auth/users/:id", handler.DeleteUser)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", "/api/v1/auth/users/"+test.userID, nil)

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedRequestBody, w.Body.String())
		})
	}
}
