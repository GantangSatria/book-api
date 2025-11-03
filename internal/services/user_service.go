package services

import (
	"book-api/internal/models"
	"book-api/internal/repository"
)

type UserService struct{ repo repository.UserRepo }

func NewUserService(r repository.UserRepo) *UserService { return &UserService{r} }

func (s *UserService) FindByUsername(username string) (models.User, error) { return s.repo.FindByUsername(username) }
func (s *UserService) Create(u models.User) (int, error) { return s.repo.Create(u) }
