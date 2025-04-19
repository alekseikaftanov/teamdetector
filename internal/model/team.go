package model

import "time"

type Team struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name" binding:"required"`
	Description string    `json:"description" db:"description"`
	CompanyID   int       `json:"company_id" db:"company_id" binding:"required"`
	CreatedBy   int       `json:"created_by" db:"created_by"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type CreateTeamInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	CompanyID   int    `json:"company_id" binding:"required"`
}

type AddUserToTeamInput struct {
	TeamID int    `json:"team_id" binding:"required"`
	Email  string `json:"email" binding:"required,email"`
	Name   string `json:"name" binding:"required"`
}

type AddUsersToTeamInput struct {
	TeamID int `json:"team_id" binding:"required"`
	Users  []struct {
		Email string `json:"email" binding:"required,email"`
		Name  string `json:"name" binding:"required"`
	} `json:"users" binding:"required,min=1"`
}
