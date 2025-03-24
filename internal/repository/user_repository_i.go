package repository

import (
	"context"
	"testovoe/internal/domain"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUser(ctx context.Context, id int64) (*domain.User, error)
	UpdateUser(ctx context.Context, id int64, user *domain.User) error
	DeleteUser(ctx context.Context, id int64) error
}
