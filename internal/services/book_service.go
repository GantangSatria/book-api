package services

import (
	"book-api/internal/models"
	"book-api/internal/repository"
	"errors"
)

type BookService struct{ repo repository.BookRepo }

func NewBookService(r repository.BookRepo) *BookService { return &BookService{r} }

func (s *BookService) Create(b models.Book) (int, error) {
	if b.ReleaseYear < 1980 || b.ReleaseYear > 2024 {
		return 0, errors.New("release_year must be between 1980 and 2024")
	}
	if b.TotalPage > 100 { b.Thickness = "tebal" } else { b.Thickness = "tipis" }

	return s.repo.Create(b)
}

func (s *BookService) GetAll() ([]models.Book, error) { return s.repo.GetAll() }
func (s *BookService) GetByID(id int) (models.Book, error) { return s.repo.GetByID(id) }
func (s *BookService) Delete(id int) error { return s.repo.Delete(id) }
func (s *BookService) GetByCategory(id int) ([]models.Book, error) { return s.repo.GetByCategory(id) }
