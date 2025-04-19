package model

import "time"

type Company struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	CreatedBy   int       `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateCompanyInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}
