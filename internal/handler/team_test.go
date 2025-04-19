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

func TestHandler_CreateTeam(t *testing.T) {
	type mockBehavior func(s *mocks.Team, team model.Team)

	testTable := []struct {
		name                string
		inputBody           string
		inputTeam           model.Team
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			inputBody: `{
				"name": "Test Team",
				"description": "Test Description",
				"company_id": 1
			}`,
			inputTeam: model.Team{
				Name:        "Test Team",
				Description: "Test Description",
				CompanyID:   1,
				CreatedBy:   1,
			},
			mockBehavior: func(s *mocks.Team, team model.Team) {
				s.On("CreateTeam", team).Return(1, nil)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name: "Empty Fields",
			inputBody: `{
				"name": "",
				"description": "",
				"company_id": 0
			}`,
			mockBehavior:        func(s *mocks.Team, team model.Team) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error":"Key: 'CreateTeamInput.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'CreateTeamInput.CompanyID' Error:Field validation for 'CompanyID' failed on the 'required' tag"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Dependencies
			c := gin.New()
			teamMock := mocks.NewTeam(t)
			testCase.mockBehavior(teamMock, testCase.inputTeam)

			services := &service.Service{Team: teamMock}
			handler := NewHandler(services)

			// Test Server
			c.POST("/api/v1/teams", func(c *gin.Context) {
				c.Set("userID", 1)
				handler.CreateTeam(c)
			})

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/teams",
				bytes.NewBufferString(testCase.inputBody))

			// Perform Request
			c.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_GetTeams(t *testing.T) {
	type mockBehavior func(s *mocks.Team, companyID int)

	testTable := []struct {
		name                string
		companyID           string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			companyID: "1",
			mockBehavior: func(s *mocks.Team, companyID int) {
				s.On("GetTeamsByCompanyID", companyID).Return([]model.Team{
					{
						ID:          1,
						Name:        "Test Team",
						Description: "Test Description",
						CompanyID:   1,
						CreatedBy:   1,
					},
				}, nil)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `[{"id":1,"name":"Test Team","description":"Test Description","company_id":1,"created_by":1,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}]`,
		},
		{
			name:                "Invalid Company ID",
			companyID:           "invalid",
			mockBehavior:        func(s *mocks.Team, companyID int) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error":"invalid company id"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Dependencies
			c := gin.New()
			teamMock := mocks.NewTeam(t)
			companyID, _ := strconv.Atoi(testCase.companyID)
			testCase.mockBehavior(teamMock, companyID)

			services := &service.Service{Team: teamMock}
			handler := NewHandler(services)

			// Test Server
			c.GET("/api/v1/teams/company/:company_id", handler.GetTeams)

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/v1/teams/company/"+testCase.companyID, nil)

			// Perform Request
			c.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_DeleteTeam(t *testing.T) {
	type mockBehavior func(s *mocks.Team, id int)

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
			mockBehavior: func(s *mocks.Team, id int) {
				s.On("DeleteTeam", id).Return(nil)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `{"message":"team deleted successfully"}`,
		},
		{
			name:                "Invalid ID",
			inputID:             "invalid",
			mockBehavior:        func(s *mocks.Team, id int) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error":"invalid id"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Dependencies
			c := gin.New()
			teamMock := mocks.NewTeam(t)
			id, _ := strconv.Atoi(testCase.inputID)
			testCase.mockBehavior(teamMock, id)

			services := &service.Service{Team: teamMock}
			handler := NewHandler(services)

			// Test Server
			c.DELETE("/api/v1/teams/team/:id", handler.DeleteTeam)

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", "/api/v1/teams/team/"+testCase.inputID, nil)

			// Perform Request
			c.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
