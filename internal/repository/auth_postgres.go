package repository

import (
	"database/sql"

	"github.com/teamdetected/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := `INSERT INTO users (email, password, name) VALUES ($1, $2, $3) RETURNING id`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	err = r.db.QueryRow(query, user.Email, string(hashedPassword), user.Name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (model.User, error) {
	var user model.User
	query := `SELECT id, email, password, name, created_at FROM users WHERE email = $1`

	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.CreatedAt)
	if err != nil {
		return model.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
