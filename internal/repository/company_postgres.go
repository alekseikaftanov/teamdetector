package repository

import (
	"database/sql"

	"github.com/teamdetected/internal/model"
)

type CompanyPostgres struct {
	db *sql.DB
}

func NewCompanyPostgres(db *sql.DB) *CompanyPostgres {
	return &CompanyPostgres{db: db}
}

func (r *CompanyPostgres) CreateCompany(company model.Company) (int, error) {
	var id int
	query := `INSERT INTO companies (name, description, created_by) VALUES ($1, $2, $3) RETURNING id`

	err := r.db.QueryRow(query, company.Name, company.Description, company.CreatedBy).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *CompanyPostgres) GetCompanyByID(id int) (model.Company, error) {
	var company model.Company
	query := `SELECT id, name, description, created_by, created_at, updated_at FROM companies WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&company.ID,
		&company.Name,
		&company.Description,
		&company.CreatedBy,
		&company.CreatedAt,
		&company.UpdatedAt,
	)
	if err != nil {
		return model.Company{}, err
	}

	return company, nil
}

func (r *CompanyPostgres) GetCompaniesByUserID(userID int) ([]model.Company, error) {
	query := `SELECT id, name, description, created_by, created_at, updated_at FROM companies WHERE created_by = $1`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies []model.Company
	for rows.Next() {
		var company model.Company
		err := rows.Scan(
			&company.ID,
			&company.Name,
			&company.Description,
			&company.CreatedBy,
			&company.CreatedAt,
			&company.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}

	return companies, nil
}

func (r *CompanyPostgres) DeleteCompany(id int) error {
	query := `DELETE FROM companies WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
