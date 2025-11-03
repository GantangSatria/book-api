package services

import (
	"book-api/internal/models"
	"book-api/internal/repository"
)

type CategoryService struct{ repo repository.CategoryRepo }

func NewCategoryService(r repository.CategoryRepo) *CategoryService { return &CategoryService{r} }

func (s *CategoryService) Create(c models.Category) (int, error) { return s.repo.Create(c) }
func (s *CategoryService) GetAll() ([]models.Category, error) { return s.repo.GetAll() }
func (s *CategoryService) GetByID(id int) (models.Category, error) { return s.repo.GetByID(id) }
func (s *CategoryService) Delete(id int) error { return s.repo.Delete(id) }
