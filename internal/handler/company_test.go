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

func TestHandler_CreateCompany(t *testing.T) {
	type mockBehavior func(s *mocks.Company, company model.Company)

	testTable := []struct {
		name                string
		inputBody           string
		inputCompany        model.Company
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			inputBody: `{
				"name": "Test Company",
				"description": "Test Description"
			}`,
			inputCompany: model.Company{
				Name:        "Test Company",
				Description: "Test Description",
				CreatedBy:   1,
			},
			mockBehavior: func(s *mocks.Company, company model.Company) {
				s.On("CreateCompany", company).Return(1, nil)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name: "Empty Fields",
			inputBody: `{
				"name": "",
				"description": ""
			}`,
			mockBehavior:        func(s *mocks.Company, company model.Company) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error":"Key: 'CreateCompanyInput.Name' Error:Field validation for 'Name' failed on the 'required' tag"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Dependencies
			c := gin.New()
			companyMock := mocks.NewCompany(t)
			testCase.mockBehavior(companyMock, testCase.inputCompany)

			services := &service.Service{Company: companyMock}
			handler := NewHandler(services)

			// Test Server
			c.POST("/api/v1/companies", func(c *gin.Context) {
				c.Set("userID", 1)
				handler.CreateCompany(c)
			})

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/companies",
				bytes.NewBufferString(testCase.inputBody))

			// Perform Request
			c.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_GetCompanies(t *testing.T) {
	type mockBehavior func(s *mocks.Company, userID int)

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
			mockBehavior: func(s *mocks.Company, userID int) {
				s.On("GetCompaniesByUserID", userID).Return([]model.Company{
					{
						ID:          1,
						Name:        "Test Company",
						Description: "Test Description",
						CreatedBy:   1,
					},
				}, nil)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `[{"id":1,"name":"Test Company","description":"Test Description","created_by":1,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}]`,
		},
		{
			name:                "Unauthorized",
			userID:              "",
			mockBehavior:        func(s *mocks.Company, userID int) {},
			expectedStatusCode:  http.StatusUnauthorized,
			expectedRequestBody: `{"error":"user not authenticated"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Dependencies
			c := gin.New()
			companyMock := mocks.NewCompany(t)
			userID, _ := strconv.Atoi(testCase.userID)
			testCase.mockBehavior(companyMock, userID)

			services := &service.Service{Company: companyMock}
			handler := NewHandler(services)

			// Test Server
			c.GET("/api/v1/companies", func(c *gin.Context) {
				if testCase.userID != "" {
					c.Set("userID", userID)
				}
				handler.GetCompanies(c)
			})

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/v1/companies", nil)

			// Perform Request
			c.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_DeleteCompany(t *testing.T) {
	type mockBehavior func(s *mocks.Company, id int)

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
			mockBehavior: func(s *mocks.Company, id int) {
				s.On("DeleteCompany", id).Return(nil)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `{"message":"company deleted successfully"}`,
		},
		{
			name:                "Invalid ID",
			inputID:             "invalid",
			mockBehavior:        func(s *mocks.Company, id int) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error":"invalid id"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Dependencies
			c := gin.New()
			companyMock := mocks.NewCompany(t)
			id, _ := strconv.Atoi(testCase.inputID)
			testCase.mockBehavior(companyMock, id)

			services := &service.Service{Company: companyMock}
			handler := NewHandler(services)

			// Test Server
			c.DELETE("/api/v1/companies/:id", handler.DeleteCompany)

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", "/api/v1/companies/"+testCase.inputID, nil)

			// Perform Request
			c.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
