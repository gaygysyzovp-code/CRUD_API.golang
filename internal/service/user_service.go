package service

import (
	"CRUD_API/internal/models"
	"CRUD_API/internal/repository"
	"context"
)

type UserService interface {
	CreateUser(ctx context.Context, req models.CreateUser) (models.User, error)
	GetUsers(ctx context.Context) ([]models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, req models.CreateUser) (models.User, error) {
	return s.repo.Create(ctx, req)
}

func (s *userService) GetUsers(ctx context.Context) ([]models.User, error) {
	return s.repo.Get(ctx)
}
