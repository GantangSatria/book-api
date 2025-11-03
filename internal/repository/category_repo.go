package repository

import (
	"book-api/internal/models"
	"database/sql"
	"errors"
)

type CategoryRepo interface {
	Create(models.Category) (int, error)
	GetAll() ([]models.Category, error)
	GetByID(int) (models.Category, error)
	Delete(int) error
}

type categoryRepo struct{ db *sql.DB }

func NewCategoryRepo(db *sql.DB) CategoryRepo { return &categoryRepo{db} }

func (r *categoryRepo) Create(c models.Category) (int, error) {
	var id int
	err := r.db.QueryRow(`INSERT INTO categories (name, created_by) VALUES ($1,$2) RETURNING id`, c.Name, c.CreatedBy).Scan(&id)
	return id, err
}

func (r *categoryRepo) GetAll() ([]models.Category, error) {
	rows, err := r.db.Query(`SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories`)
	if err != nil { return nil, err }
	defer rows.Close()

	var list []models.Category
	for rows.Next() {
		var c models.Category
		rows.Scan(&c.ID, &c.Name, &c.CreatedAt, &c.CreatedBy, &c.ModifiedAt, &c.ModifiedBy)
		list = append(list, c)
	}
	return list, nil
}

func (r *categoryRepo) GetByID(id int) (models.Category, error) {
	var c models.Category
	err := r.db.QueryRow(`SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories WHERE id=$1`, id).
		Scan(&c.ID,&c.Name,&c.CreatedAt,&c.CreatedBy,&c.ModifiedAt,&c.ModifiedBy)
	if err == sql.ErrNoRows { return c, errors.New("category not found") }
	return c, err
}

func (r *categoryRepo) Delete(id int) error {
	res, err := r.db.Exec(`DELETE FROM categories WHERE id=$1`, id)
	if err != nil { return err }
	n, _ := res.RowsAffected()
	if n == 0 { return errors.New("category not found") }
	return nil
}
