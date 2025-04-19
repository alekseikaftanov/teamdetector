package repository

import (
	"database/sql"

	"github.com/teamdetected/internal/model"
)

type SurveyPostgres struct {
	db *sql.DB
}

func NewSurveyPostgres(db *sql.DB) *SurveyPostgres {
	return &SurveyPostgres{db: db}
}

func (r *SurveyPostgres) CreateSurvey(survey model.Survey) (int, error) {
	var id int
	query := `INSERT INTO surveys (team_id, status, created_by) 
              VALUES ($1, $2, $3) RETURNING id`

	err := r.db.QueryRow(query, survey.TeamID, survey.Status, survey.CreatedBy).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *SurveyPostgres) GetSurveyByID(id int) (model.Survey, error) {
	var survey model.Survey
	query := `SELECT id, team_id, status, created_by, created_at, updated_at 
              FROM surveys WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&survey.ID, &survey.TeamID, &survey.Status, &survey.CreatedBy,
		&survey.CreatedAt, &survey.UpdatedAt,
	)
	if err != nil {
		return model.Survey{}, err
	}

	return survey, nil
}

func (r *SurveyPostgres) GetSurveysByTeamID(teamID int) ([]model.Survey, error) {
	query := `SELECT id, team_id, status, created_by, created_at, updated_at 
              FROM surveys WHERE team_id = $1`

	rows, err := r.db.Query(query, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var surveys []model.Survey
	for rows.Next() {
		var survey model.Survey
		err := rows.Scan(
			&survey.ID, &survey.TeamID, &survey.Status, &survey.CreatedBy,
			&survey.CreatedAt, &survey.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		surveys = append(surveys, survey)
	}

	return surveys, nil
}

func (r *SurveyPostgres) DeleteSurvey(id int) error {
	query := `DELETE FROM surveys WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *SurveyPostgres) CreateSurveyResponse(response model.SurveyResponse) (int, error) {
	var id int
	query := `INSERT INTO survey_responses (survey_id, user_id, question_id, option_id) 
              VALUES ($1, $2, $3, $4) RETURNING id`

	err := r.db.QueryRow(query, response.SurveyID, response.UserID, response.QuestionID, response.OptionID).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *SurveyPostgres) GetSurveyResponses(surveyID int) ([]model.SurveyResponse, error) {
	query := `SELECT id, survey_id, user_id, question_id, option_id, created_at 
              FROM survey_responses WHERE survey_id = $1`

	rows, err := r.db.Query(query, surveyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var responses []model.SurveyResponse
	for rows.Next() {
		var response model.SurveyResponse
		err := rows.Scan(
			&response.ID, &response.SurveyID, &response.UserID, &response.QuestionID,
			&response.OptionID, &response.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (r *SurveyPostgres) GetSurveyOptions() ([]model.SurveyOption, error) {
	query := `SELECT id, text, value FROM survey_options`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var options []model.SurveyOption
	for rows.Next() {
		var option model.SurveyOption
		err := rows.Scan(&option.ID, &option.Text, &option.Value)
		if err != nil {
			return nil, err
		}
		options = append(options, option)
	}

	return options, nil
}

func (r *SurveyPostgres) GetSurveyQuestions() ([]model.SurveyQuestion, error) {
	query := `SELECT id, text, category FROM survey_questions`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []model.SurveyQuestion
	for rows.Next() {
		var question model.SurveyQuestion
		err := rows.Scan(&question.ID, &question.Text, &question.Category)
		if err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}

	return questions, nil
}
