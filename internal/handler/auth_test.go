package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
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
			name: "OK",
			inputBody: `{
				"email": "test@test.com",
				"password": "test123456",
				"name": "Test User",
				"role": "user"
			}`,
			inputUser: model.User{
				Email:    "test@test.com",
				Password: "test123456",
				Name:     "Test User",
				Role:     "user",
			},
			mockBehavior: func(s *mocks.Authorization, user model.User) {
				s.On("CreateUser", user).Return(1, nil)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name: "Empty Fields",
			inputBody: `{
				"email": "",
				"password": "",
				"name": "",
				"role": ""
			}`,
			mockBehavior:        func(s *mocks.Authorization, user model.User) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error":"Key: 'SignUpInput.Email' Error:Field validation for 'Email' failed on the 'required' tag\nKey: 'SignUpInput.Password' Error:Field validation for 'Password' failed on the 'required' tag\nKey: 'SignUpInput.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'SignUpInput.Role' Error:Field validation for 'Role' failed on the 'required' tag"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Dependencies
			c := gin.New()
			authMock := mocks.NewAuthorization(t)
			testCase.mockBehavior(authMock, testCase.inputUser)

			services := &service.Service{Authorization: authMock}
			handler := NewHandler(services)

			// Test Server
			c.POST("/api/v1/auth/register", handler.Register)

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/auth/register",
				bytes.NewBufferString(testCase.inputBody))

			// Perform Request
			c.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_Login(t *testing.T) {
	type mockBehavior func(s *mocks.Authorization, email, password string)

	testTable := []struct {
		name                string
		inputBody           string
		email               string
		password            string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			inputBody: `{
				"email": "test@test.com",
				"password": "test123"
			}`,
			email:    "test@test.com",
			password: "test123",
			mockBehavior: func(s *mocks.Authorization, email, password string) {
				s.On("GenerateToken", email, password).Return("test-token", nil)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `{"token":"test-token"}`,
		},
		{
			name: "Wrong Input",
			inputBody: `{
				"email": "test@test.com"
			}`,
			mockBehavior:        func(s *mocks.Authorization, email, password string) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error":"Key: 'SignInInput.Password' Error:Field validation for 'Password' failed on the 'required' tag"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Dependencies
			c := gin.New()
			authMock := mocks.NewAuthorization(t)
			testCase.mockBehavior(authMock, testCase.email, testCase.password)

			services := &service.Service{Authorization: authMock}
			handler := NewHandler(services)

			// Test Server
			c.POST("/api/v1/auth/login", handler.Login)

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/auth/login",
				bytes.NewBufferString(testCase.inputBody))

			// Perform Request
			c.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_DeleteUser(t *testing.T) {
	type mockBehavior func(s *mocks.Authorization, id int)

	testTable := []struct {
		name                string
		inputID             string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:    "OK",
			inputID: "1",
			mockBehavior: func(s *mocks.Authorization, id int) {
				s.On("DeleteUser", id).Return(nil)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `{"message":"user deleted successfully"}`,
		},
		{
			name:                "Invalid ID",
			inputID:             "invalid",
			mockBehavior:        func(s *mocks.Authorization, id int) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error":"invalid id"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Dependencies
			c := gin.New()
			authMock := mocks.NewAuthorization(t)
			id, _ := strconv.Atoi(testCase.inputID)
			testCase.mockBehavior(authMock, id)

			services := &service.Service{Authorization: authMock}
			handler := NewHandler(services)

			// Test Server
			c.DELETE("/api/v1/auth/users/:id", handler.DeleteUser)

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", "/api/v1/auth/users/"+testCase.inputID, nil)

			// Perform Request
			c.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
