package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teamdetected/internal/model"
)

func (h *Handler) CreateSurvey(c *gin.Context) {
	var input model.CreateSurveyInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	survey := model.Survey{
		TeamID:    input.TeamID,
		CreatedBy: userID.(int),
	}

	id, err := h.services.Survey.CreateSurvey(survey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) GetSurvey(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("survey_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid survey id"})
		return
	}

	survey, err := h.services.Survey.GetSurveyByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, survey)
}

func (h *Handler) GetSurveysByTeam(c *gin.Context) {
	teamID, err := strconv.Atoi(c.Param("team_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid team id"})
		return
	}

	surveys, err := h.services.Survey.GetSurveysByTeamID(teamID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, surveys)
}

func (h *Handler) DeleteSurvey(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("survey_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid survey id"})
		return
	}

	err = h.services.Survey.DeleteSurvey(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "survey deleted successfully"})
}

func (h *Handler) CreateSurveyResponse(c *gin.Context) {
	var input model.CreateSurveyResponseInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	response := model.SurveyResponse{
		SurveyID:   input.SurveyID,
		UserID:     userID.(int),
		QuestionID: input.QuestionID,
		OptionID:   input.OptionID,
	}

	id, err := h.services.Survey.CreateSurveyResponse(response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) GetSurveyResponses(c *gin.Context) {
	surveyID, err := strconv.Atoi(c.Param("survey_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid survey id"})
		return
	}

	responses, err := h.services.Survey.GetSurveyResponses(surveyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses)
}

func (h *Handler) GetSurveyOptions(c *gin.Context) {
	options, err := h.services.Survey.GetSurveyOptions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, options)
}

func (h *Handler) GetSurveyQuestions(c *gin.Context) {
	questions, err := h.services.Survey.GetSurveyQuestions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, questions)
}
