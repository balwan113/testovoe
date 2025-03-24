package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"testovoe/internal/domain"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"

	if err := r.db.QueryRow(ctx, query, user.Name, user.Email).Scan(&user.ID); err != nil {
		return fmt.Errorf("ошибка при создании пользователя: %w", err)
	}
	return nil
}

func (r *UserRepository) GetUser(ctx context.Context, id int64) (*domain.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = $1"

	var user domain.User
	if err := r.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("пользователь не найден")
		}
		return nil, fmt.Errorf("ошибка при получении пользователя: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, id int64, user *domain.User) error {
	query := "UPDATE users SET name = $1, email = $2 WHERE id = $3"
	cmdTag, err := r.db.Exec(ctx, query, user.Name, user.Email, id)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении пользователя с id %d: %w", id, err)
	}
	if cmdTag.RowsAffected() == 0 {
		return errors.New("пользователь не найден")
	}
	return nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int64) error {
	query := "DELETE FROM users WHERE id = $1"
	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("ошибка при удалении пользователя с id %d: %w", id, err)
	}
	if cmdTag.RowsAffected() == 0 {
		return errors.New("пользователь не найден")
	}
	return nil
}
