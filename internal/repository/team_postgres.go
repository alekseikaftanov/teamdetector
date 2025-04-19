package repository

import (
	"database/sql"

	"github.com/teamdetected/internal/model"
)

type TeamPostgres struct {
	db *sql.DB
}

func NewTeamPostgres(db *sql.DB) *TeamPostgres {
	return &TeamPostgres{db: db}
}

func (r *TeamPostgres) CreateTeam(team model.Team) (int, error) {
	var id int
	query := `INSERT INTO teams (name, description, company_id, created_by) VALUES ($1, $2, $3, $4) RETURNING id`

	err := r.db.QueryRow(query, team.Name, team.Description, team.CompanyID, team.CreatedBy).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *TeamPostgres) GetTeamByID(id int) (model.Team, error) {
	var team model.Team
	query := `SELECT id, name, description, company_id, created_by, created_at, updated_at FROM teams WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&team.ID,
		&team.Name,
		&team.Description,
		&team.CompanyID,
		&team.CreatedBy,
		&team.CreatedAt,
		&team.UpdatedAt,
	)
	if err != nil {
		return model.Team{}, err
	}

	return team, nil
}

func (r *TeamPostgres) GetTeamsByCompanyID(companyID int) ([]model.Team, error) {
	query := `SELECT id, name, description, company_id, created_by, created_at, updated_at FROM teams WHERE company_id = $1`
	rows, err := r.db.Query(query, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []model.Team
	for rows.Next() {
		var team model.Team
		err := rows.Scan(
			&team.ID,
			&team.Name,
			&team.Description,
			&team.CompanyID,
			&team.CreatedBy,
			&team.CreatedAt,
			&team.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}

	return teams, nil
}

func (r *TeamPostgres) DeleteTeam(id int) error {
	query := `DELETE FROM teams WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
