package repository

import (
	"context"
	"testovoe/internal/models"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, id int64) (*models.User, error)
	UpdateUser(ctx context.Context, id int64, user *models.User) error
	DeleteUser(ctx context.Context, id int64) error
}