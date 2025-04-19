package model

import "time"

type User struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" binding:"required" db:"name"`
	Email     string    `json:"email" binding:"required,email" db:"email"`
	Password  string    `json:"password" binding:"required" db:"password_hash"`
	Role      string    `json:"role" binding:"required" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type SignUpInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Role  string `json:"role" binding:"required"`
}

type UserRole string

const (
	UserRoleAdmin   UserRole = "admin"
	UserRoleTeam    UserRole = "team"
	UserRoleManager UserRole = "manager"
)
