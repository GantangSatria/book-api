package repository

import (
	"book-api/internal/models"
	"database/sql"
	"errors"
)

type UserRepo interface {
	FindByUsername(string) (models.User, error)
	Create(models.User) (int, error)
}

type userRepo struct{ db *sql.DB }

func NewUserRepo(db *sql.DB) UserRepo { return &userRepo{db} }

func (r *userRepo) FindByUsername(username string) (models.User, error) {
	var u models.User
	err := r.db.QueryRow(`SELECT id, username, password, created_at FROM users WHERE username=$1`, username).
		Scan(&u.ID,&u.Username,&u.Password,&u.CreatedAt)
	if err == sql.ErrNoRows { return u, errors.New("user not found") }
	return u, err
}

func (r *userRepo) Create(u models.User) (int, error) {
	query := `INSERT INTO users (username, password, created_at, created_by)
	          VALUES ($1, $2, NOW(), $1) RETURNING id`
	var id int
	err := r.db.QueryRow(query, u.Username, u.Password).Scan(&id)
	return id, err
}

