package service

import (
	"context"
	"errors"
	"testovoe/internal/models"
	"testovoe/internal/repository"
)





type UserService struct {
	repo repository.UserRepositoryInterface
}


func NewUserService(repo repository.UserRepositoryInterface) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("Имя или email не могут быть пустыми")
	}

	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) GetUser(ctx context.Context, id int64) (*models.User, error) {
	return s.repo.GetUser(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, id int64, user *models.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("Имя или email не могут быть пустыми")
	}

	return s.repo.UpdateUser(ctx, id, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.DeleteUser(ctx, id)
}
