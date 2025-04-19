package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/teamdetected/internal/model"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := `INSERT INTO users (name, email, password_hash) VALUES (:name, :email, :password) RETURNING id`
	rows, err := r.db.NamedQuery(query, map[string]interface{}{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
	})
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

func (r *AuthPostgres) GetUser(email, password string) (model.User, error) {
	var user model.User
	query := `SELECT id, name, email, password_hash FROM users WHERE email = $1`
	err := r.db.Get(&user, query, email)
	return user, err
}

func (r *AuthPostgres) GetUserByID(id int) (model.User, error) {
	var user model.User
	query := `SELECT id, name, email, password_hash FROM users WHERE id = $1`
	err := r.db.Get(&user, query, id)
	return user, err
}

func (r *AuthPostgres) UpdateUser(id int, user model.User) error {
	query := `UPDATE users SET name = :name, email = :email WHERE id = :id`
	params := map[string]interface{}{
		"id":    id,
		"name":  user.Name,
		"email": user.Email,
	}
	result, err := r.db.NamedExec(query, params)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("user with id %d not found", id)
	}
	return nil
}

func (r *AuthPostgres) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = :id`
	result, err := r.db.NamedExec(query, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("user with id %d not found", id)
	}
	return nil
}

func (r *AuthPostgres) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	query := `SELECT id, email, name, role FROM users WHERE email = $1`
	err := r.db.Get(&user, query, email)
	return user, err
}
