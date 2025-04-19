package model

import "time"

type Survey struct {
	ID        int       `json:"id"`
	TeamID    int       `json:"team_id" binding:"required"`
	Status    string    `json:"status"` // active, completed
	CreatedBy int       `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SurveyQuestion struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	Category string `json:"category"`
}

type SurveyResponse struct {
	ID         int       `json:"id"`
	SurveyID   int       `json:"survey_id" binding:"required"`
	UserID     int       `json:"user_id" binding:"required"`
	QuestionID int       `json:"question_id" binding:"required"`
	OptionID   int       `json:"option_id" binding:"required"`
	CreatedAt  time.Time `json:"created_at"`
}

type SurveyOption struct {
	ID    int    `json:"id"`
	Text  string `json:"text"`
	Value int    `json:"value"`
}

type CreateSurveyInput struct {
	TeamID int `json:"team_id" binding:"required"`
}

type CreateSurveyResponseInput struct {
	SurveyID   int `json:"survey_id" binding:"required"`
	QuestionID int `json:"question_id" binding:"required"`
	OptionID   int `json:"option_id" binding:"required"`
}
