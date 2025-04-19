package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/teamdetected/internal/model"
)

type SurveyPostgres struct {
	db *sqlx.DB
}

func NewSurveyPostgres(db *sqlx.DB) *SurveyPostgres {
	return &SurveyPostgres{db: db}
}

func (r *SurveyPostgres) CreateSurvey(survey model.Survey) (int, error) {
	var id int
	query := `INSERT INTO surveys (team_id, status, created_by) 
              VALUES (:team_id, :status, :created_by) RETURNING id`

	params := map[string]interface{}{
		"team_id":    survey.TeamID,
		"status":     survey.Status,
		"created_by": survey.CreatedBy,
	}

	rows, err := r.db.NamedQuery(query, params)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (r *SurveyPostgres) GetSurveyByID(id int) (model.Survey, error) {
	var survey model.Survey
	query := `SELECT id, team_id, status, created_by, created_at, updated_at 
              FROM surveys WHERE id = $1`

	err := r.db.Get(&survey, query, id)
	if err != nil {
		return model.Survey{}, err
	}

	return survey, nil
}

func (r *SurveyPostgres) GetSurveysByTeamID(teamID int) ([]model.Survey, error) {
	var surveys []model.Survey
	query := `SELECT id, team_id, status, created_by, created_at, updated_at 
              FROM surveys WHERE team_id = $1`

	err := r.db.Select(&surveys, query, teamID)
	if err != nil {
		return nil, err
	}

	return surveys, nil
}

func (r *SurveyPostgres) DeleteSurvey(id int) error {
	query := `DELETE FROM surveys WHERE id = :id`
	result, err := r.db.NamedExec(query, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("survey with id %d not found", id)
	}
	return nil
}

func (r *SurveyPostgres) CreateSurveyResponse(response model.SurveyResponse) (int, error) {
	var id int
	query := `INSERT INTO survey_responses (survey_id, user_id, question_id, option_id) 
              VALUES (:survey_id, :user_id, :question_id, :option_id) RETURNING id`

	params := map[string]interface{}{
		"survey_id":   response.SurveyID,
		"user_id":     response.UserID,
		"question_id": response.QuestionID,
		"option_id":   response.OptionID,
	}

	rows, err := r.db.NamedQuery(query, params)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (r *SurveyPostgres) GetSurveyResponses(surveyID int) ([]model.SurveyResponse, error) {
	var responses []model.SurveyResponse
	query := `SELECT id, survey_id, user_id, question_id, option_id, created_at 
              FROM survey_responses WHERE survey_id = $1`

	err := r.db.Select(&responses, query, surveyID)
	if err != nil {
		return nil, err
	}

	return responses, nil
}

func (r *SurveyPostgres) GetSurveyOptions() ([]model.SurveyOption, error) {
	var options []model.SurveyOption
	query := `SELECT id, text, value FROM survey_options`

	err := r.db.Select(&options, query)
	if err != nil {
		return nil, err
	}

	return options, nil
}

func (r *SurveyPostgres) GetSurveyQuestions() ([]model.SurveyQuestion, error) {
	var questions []model.SurveyQuestion
	query := `SELECT id, text, category FROM survey_questions`

	err := r.db.Select(&questions, query)
	if err != nil {
		return nil, err
	}

	return questions, nil
}
