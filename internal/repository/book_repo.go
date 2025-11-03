package repository

import (
	"book-api/internal/models"
	"database/sql"
	"errors"
)

type BookRepo interface {
	Create(models.Book) (int, error)
	GetAll() ([]models.Book, error)
	GetByID(int) (models.Book, error)
	Delete(int) error
	GetByCategory(int) ([]models.Book, error)
}

type bookRepo struct{ db *sql.DB }

func NewBookRepo(db *sql.DB) BookRepo { return &bookRepo{db} }

func (r *bookRepo) Create(b models.Book) (int, error) {
	var id int
	err := r.db.QueryRow(`INSERT INTO books
	(title, description, image_url, release_year, price, total_page, thickness, category_id, created_by)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id`,
		b.Title,b.Description,b.ImageURL,b.ReleaseYear,b.Price,b.TotalPage,b.Thickness,b.CategoryID,b.CreatedBy).Scan(&id)
	return id, err
}

func (r *bookRepo) GetAll() ([]models.Book, error) {
	rows, err := r.db.Query(`SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM books`)
	if err != nil { return nil, err }
	defer rows.Close()

	var list []models.Book
	for rows.Next() {
		var b models.Book
		rows.Scan(&b.ID,&b.Title,&b.Description,&b.ImageURL,&b.ReleaseYear,&b.Price,&b.TotalPage,&b.Thickness,&b.CategoryID,&b.CreatedAt,&b.CreatedBy,&b.ModifiedAt,&b.ModifiedBy)
		list = append(list, b)
	}
	return list, nil
}

func (r *bookRepo) GetByID(id int) (models.Book, error) {
	var b models.Book
	err := r.db.QueryRow(`SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM books WHERE id=$1`, id).
		Scan(&b.ID,&b.Title,&b.Description,&b.ImageURL,&b.ReleaseYear,&b.Price,&b.TotalPage,&b.Thickness,&b.CategoryID,&b.CreatedAt,&b.CreatedBy,&b.ModifiedAt,&b.ModifiedBy)
	if err == sql.ErrNoRows { return b, errors.New("book not found") }
	return b, err
}

func (r *bookRepo) Delete(id int) error {
	res, err := r.db.Exec(`DELETE FROM books WHERE id=$1`, id)
	if err != nil { return err }
	n, _ := res.RowsAffected()
	if n == 0 { return errors.New("book not found") }
	return nil
}

func (r *bookRepo) GetByCategory(catID int) ([]models.Book, error) {
	rows, err := r.db.Query(`SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM books WHERE category_id=$1`, catID)
	if err != nil { return nil, err }
	defer rows.Close()

	var list []models.Book
	for rows.Next() {
		var b models.Book
		rows.Scan(&b.ID,&b.Title,&b.Description,&b.ImageURL,&b.ReleaseYear,&b.Price,&b.TotalPage,&b.Thickness,&b.CategoryID,&b.CreatedAt,&b.CreatedBy,&b.ModifiedAt,&b.ModifiedBy)
		list = append(list, b)
	}
	return list, nil
}
