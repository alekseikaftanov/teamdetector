package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/teamdetected/internal/model"
)

type TeamPostgres struct {
	db *sqlx.DB
}

func NewTeamPostgres(db *sqlx.DB) *TeamPostgres {
	return &TeamPostgres{db: db}
}

func (r *TeamPostgres) CreateTeam(team model.Team) (int, error) {
	var id int
	query := `INSERT INTO teams (name, description, company_id, created_by) 
              VALUES (:name, :description, :company_id, :created_by) RETURNING id`

	params := map[string]interface{}{
		"name":        team.Name,
		"description": team.Description,
		"company_id":  team.CompanyID,
		"created_by":  team.CreatedBy,
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

func (r *TeamPostgres) GetTeamByID(id int) (model.Team, error) {
	var team model.Team
	query := `SELECT id, name, description, company_id, created_by, created_at, updated_at 
              FROM teams WHERE id = $1`

	err := r.db.Get(&team, query, id)
	if err != nil {
		return model.Team{}, err
	}

	return team, nil
}

func (r *TeamPostgres) GetTeamsByCompanyID(companyID int) ([]model.Team, error) {
	var teams []model.Team
	query := `SELECT id, name, description, company_id, created_by, created_at, updated_at 
              FROM teams WHERE company_id = $1`

	err := r.db.Select(&teams, query, companyID)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func (r *TeamPostgres) DeleteTeam(id int) error {
	query := `DELETE FROM teams WHERE id = :id`
	result, err := r.db.NamedExec(query, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("team with id %d not found", id)
	}
	return nil
}
