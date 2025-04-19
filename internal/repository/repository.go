package repository

import (
	"database/sql"

	"github.com/teamdetected/internal/model"
)

type Repository struct {
	Authorization
	Company
	Team
	Survey
}

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(email, password string) (model.User, error)
	DeleteUser(id int) error
}

type Company interface {
	CreateCompany(company model.Company) (int, error)
	GetCompanyByID(id int) (model.Company, error)
	GetCompaniesByUserID(userID int) ([]model.Company, error)
	DeleteCompany(id int) error
}

type Team interface {
	CreateTeam(team model.Team) (int, error)
	GetTeamByID(id int) (model.Team, error)
	GetTeamsByCompanyID(companyID int) ([]model.Team, error)
	DeleteTeam(id int) error
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Company:       NewCompanyPostgres(db),
		Team:          NewTeamPostgres(db),
		Survey:        NewSurveyPostgres(db),
	}
}
