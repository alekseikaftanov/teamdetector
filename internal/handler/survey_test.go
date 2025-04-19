package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/teamdetected/internal/model"
	"github.com/teamdetected/internal/service"
	"github.com/teamdetected/internal/service/mocks"
)

func TestHandler_CreateSurvey(t *testing.T) {
	type mockBehavior func(s *mocks.Survey, survey model.Survey)

	testTable := []struct {
		name                string
		inputBody           string
		inputSurvey         model.Survey
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			inputBody: `{
				"team_id": 1
			}`,
			inputSurvey: model.Survey{
				TeamID:    1,
				CreatedBy: 1,
			},
			mockBehavior: func(s *mocks.Survey, survey model.Survey) {
				s.On("CreateSurvey", survey).Return(1, nil)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name: "Empty Fields",
			inputBody: `{
				"team_id": 0
			}`,
			mockBehavior: func(s *mocks.Survey, survey model.Survey) {
			},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error":"Key: 'CreateSurveyInput.TeamID' Error:Field validation for 'TeamID' failed on the 'required' tag"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			survey := new(mocks.Survey)
			testCase.mockBehavior(survey, testCase.inputSurvey)

			services := &service.Service{Survey: survey}
			handler := NewHandler(services)

			// Init Endpoint
			r := gin.New()
			r.POST("/api/v1/surveys", func(c *gin.Context) {
				c.Set("userID", 1)
				handler.CreateSurvey(c)
			})

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/surveys",
				bytes.NewBufferString(testCase.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_GetSurvey(t *testing.T) {
	type mockBehavior func(s *mocks.Survey, id int)

	testTable := []struct {
		name                string
		inputId             string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:    "OK",
			inputId: "1",
			mockBehavior: func(s *mocks.Survey, id int) {
				s.On("GetSurveyByID", id).Return(model.Survey{
					ID:        1,
					TeamID:    1,
					Status:    "active",
					CreatedBy: 1,
				}, nil)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `{"id":1,"team_id":1,"status":"active","created_by":1,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`,
		},
		{
			name:                "Invalid ID",
			inputId:             "invalid",
			mockBehavior:        func(s *mocks.Survey, id int) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error":"invalid survey id"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			survey := new(mocks.Survey)
			id, _ := strconv.Atoi(testCase.inputId)
			testCase.mockBehavior(survey, id)

			services := &service.Service{Survey: survey}
			handler := NewHandler(services)

			// Init Endpoint
			r := gin.New()
			r.GET("/api/v1/surveys/:survey_id", handler.GetSurvey)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/v1/surveys/"+testCase.inputId, nil)

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_GetSurveysByTeam(t *testing.T) {
	type mockBehavior func(s *mocks.Survey, teamID int)

	testTable := []struct {
		name                string
		inputTeamID         string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:        "OK",
			inputTeamID: "1",
			mockBehavior: func(s *mocks.Survey, teamID int) {
				s.On("GetSurveysByTeamID", teamID).Return([]model.Survey{
					{
						ID:        1,
						TeamID:    1,
						Status:    "active",
						CreatedBy: 1,
					},
				}, nil)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `[{"id":1,"team_id":1,"status":"active","created_by":1,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}]`,
		},
		{
			name:                "Invalid Team ID",
			inputTeamID:         "invalid",
			mockBehavior:        func(s *mocks.Survey, teamID int) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error":"invalid team id"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			survey := new(mocks.Survey)
			teamID, _ := strconv.Atoi(testCase.inputTeamID)
			testCase.mockBehavior(survey, teamID)

			services := &service.Service{Survey: survey}
			handler := NewHandler(services)

			// Init Endpoint
			r := gin.New()
			r.GET("/api/v1/surveys/team/:team_id", handler.GetSurveysByTeam)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/v1/surveys/team/"+testCase.inputTeamID, nil)

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_DeleteSurvey(t *testing.T) {
	type mockBehavior func(s *mocks.Survey, id int)

	testTable := []struct {
		name                string
		inputId             string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:    "OK",
			inputId: "1",
			mockBehavior: func(s *mocks.Survey, id int) {
				s.On("DeleteSurvey", id).Return(nil)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `{"message":"survey deleted successfully"}`,
		},
		{
			name:                "Invalid ID",
			inputId:             "invalid",
			mockBehavior:        func(s *mocks.Survey, id int) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error":"invalid survey id"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			survey := new(mocks.Survey)
			id, _ := strconv.Atoi(testCase.inputId)
			testCase.mockBehavior(survey, id)

			services := &service.Service{Survey: survey}
			handler := NewHandler(services)

			// Init Endpoint
			r := gin.New()
			r.DELETE("/api/v1/surveys/:survey_id", handler.DeleteSurvey)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", "/api/v1/surveys/"+testCase.inputId, nil)

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_CreateSurveyResponse(t *testing.T) {
	type mockBehavior func(s *mocks.Survey, response model.SurveyResponse)

	testTable := []struct {
		name                string
		inputBody           string
		inputResponse       model.SurveyResponse
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			inputBody: `{
				"survey_id": 1,
				"question_id": 1,
				"option_id": 4
			}`,
			inputResponse: model.SurveyResponse{
				SurveyID:   1,
				UserID:     1,
				QuestionID: 1,
				OptionID:   4,
			},
			mockBehavior: func(s *mocks.Survey, response model.SurveyResponse) {
				s.On("CreateSurveyResponse", response).Return(1, nil)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name: "Empty Fields",
			inputBody: `{
				"survey_id": 0,
				"question_id": 0,
				"option_id": 0
			}`,
			mockBehavior: func(s *mocks.Survey, response model.SurveyResponse) {
			},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error":"Key: 'CreateSurveyResponseInput.SurveyID' Error:Field validation for 'SurveyID' failed on the 'required' tag\nKey: 'CreateSurveyResponseInput.QuestionID' Error:Field validation for 'QuestionID' failed on the 'required' tag\nKey: 'CreateSurveyResponseInput.OptionID' Error:Field validation for 'OptionID' failed on the 'required' tag"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			survey := new(mocks.Survey)
			testCase.mockBehavior(survey, testCase.inputResponse)

			services := &service.Service{Survey: survey}
			handler := NewHandler(services)

			// Init Endpoint
			r := gin.New()
			r.POST("/api/v1/surveys/:survey_id/responses", func(c *gin.Context) {
				c.Set("userID", 1)
				handler.CreateSurveyResponse(c)
			})

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/surveys/1/responses",
				bytes.NewBufferString(testCase.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_GetSurveyResponses(t *testing.T) {
	type mockBehavior func(s *mocks.Survey, surveyID int)

	testTable := []struct {
		name                string
		paramSurveyID       string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:          "OK",
			paramSurveyID: "1",
			mockBehavior: func(s *mocks.Survey, surveyID int) {
				responses := []model.SurveyResponse{
					{
						ID:         1,
						SurveyID:   1,
						UserID:     1,
						QuestionID: 1,
						OptionID:   4,
						CreatedAt:  time.Date(2024, 3, 20, 0, 0, 0, 0, time.UTC),
					},
				}
				s.On("GetSurveyResponses", surveyID).Return(responses, nil)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `[{"id":1,"survey_id":1,"user_id":1,"question_id":1,"option_id":4,"created_at":"2024-03-20T00:00:00Z"}]`,
		},
		{
			name:          "Invalid Survey ID",
			paramSurveyID: "invalid",
			mockBehavior: func(s *mocks.Survey, surveyID int) {
				// Не нужно устанавливать ожидание для мока, так как до него не дойдет выполнение
			},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"error":"invalid survey id"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			survey := new(mocks.Survey)
			surveyID, _ := strconv.Atoi(testCase.paramSurveyID)
			testCase.mockBehavior(survey, surveyID)

			services := &service.Service{Survey: survey}
			handler := NewHandler(services)

			// Init Endpoint
			r := gin.New()
			r.GET("/api/v1/surveys/:survey_id/responses", handler.GetSurveyResponses)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/surveys/%s/responses", testCase.paramSurveyID), nil)

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_GetSurveyOptions(t *testing.T) {
	type mockBehavior func(s *mocks.Survey)

	testTable := []struct {
		name                string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			mockBehavior: func(s *mocks.Survey) {
				s.On("GetSurveyOptions").Return([]model.SurveyOption{
					{
						ID:    1,
						Text:  "Strongly Disagree",
						Value: 1,
					},
					{
						ID:    2,
						Text:  "Disagree",
						Value: 2,
					},
				}, nil)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `[{"id":1,"text":"Strongly Disagree","value":1},{"id":2,"text":"Disagree","value":2}]`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			survey := new(mocks.Survey)
			testCase.mockBehavior(survey)

			services := &service.Service{Survey: survey}
			handler := NewHandler(services)

			// Init Endpoint
			r := gin.New()
			r.GET("/api/v1/surveys/options", handler.GetSurveyOptions)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/v1/surveys/options", nil)

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_GetSurveyQuestions(t *testing.T) {
	type mockBehavior func(s *mocks.Survey)

	testTable := []struct {
		name                string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			mockBehavior: func(s *mocks.Survey) {
				s.On("GetSurveyQuestions").Return([]model.SurveyQuestion{
					{
						ID:       1,
						Text:     "How satisfied are you with the team's communication?",
						Category: "Communication",
					},
					{
						ID:       2,
						Text:     "How well do you understand your role and responsibilities?",
						Category: "Role Clarity",
					},
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedRequestBody: `[{"id":1,"text":"How satisfied are you with the team's communication?","category":"Communication"},` +
				`{"id":2,"text":"How well do you understand your role and responsibilities?","category":"Role Clarity"}]`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			survey := new(mocks.Survey)
			testCase.mockBehavior(survey)

			services := &service.Service{Survey: survey}
			handler := NewHandler(services)

			// Init Endpoint
			r := gin.New()
			r.GET("/api/v1/surveys/questions", handler.GetSurveyQuestions)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/v1/surveys/questions", nil)

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
