package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/teamdetected/internal/model"
)

type CompanyPostgres struct {
	db *sqlx.DB
}

func NewCompanyPostgres(db *sqlx.DB) *CompanyPostgres {
	return &CompanyPostgres{db: db}
}

func (r *CompanyPostgres) CreateCompany(company model.Company) (int, error) {
	var id int
	query := `INSERT INTO companies (name, description, created_by) 
              VALUES (:name, :description, :created_by) RETURNING id`

	params := map[string]interface{}{
		"name":        company.Name,
		"description": company.Description,
		"created_by":  company.CreatedBy,
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

func (r *CompanyPostgres) GetCompanyByID(id int) (model.Company, error) {
	var company model.Company
	query := `SELECT id, name, description, created_by, created_at, updated_at 
              FROM companies WHERE id = $1`

	err := r.db.Get(&company, query, id)
	if err != nil {
		return model.Company{}, err
	}

	return company, nil
}

func (r *CompanyPostgres) GetCompaniesByUserID(userID int) ([]model.Company, error) {
	var companies []model.Company
	query := `SELECT id, name, description, created_by, created_at, updated_at 
              FROM companies WHERE created_by = $1`

	err := r.db.Select(&companies, query, userID)
	if err != nil {
		return nil, err
	}

	return companies, nil
}

func (r *CompanyPostgres) DeleteCompany(id int) error {
	query := `DELETE FROM companies WHERE id = :id`
	result, err := r.db.NamedExec(query, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("company with id %d not found", id)
	}
	return nil
}
