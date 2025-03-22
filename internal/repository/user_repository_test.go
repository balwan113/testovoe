package repository

import (
	"context"
	"errors"
	"strings"
	"testing"
	"testovoe/internal/models"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func setupDBConfig(t *testing.T) (*pgxpool.Pool, func()) {
	ctx := context.Background()

	// Запуск контейнера с PostgreSQL
	pgContainer, err := postgres.Run(ctx, "postgres:17-alpine",
		postgres.WithDatabase("test_db"),
		postgres.WithUsername("user"),
		postgres.WithPassword("password"),
		postgres.WithSQLDriver("pgx"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(30*time.Second),
		),
	)
	if err != nil {
		t.Fatalf("Не удалось запустить контейнер с PostgreSQL: %v", err)
	}

	dsn, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("Не удалось получить строку подключения: %v", err)
	}
	t.Logf("Строка подключения: %s", dsn)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		t.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	_, err = pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			id BIGSERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(100) UNIQUE NOT NULL
		)
	`)
	if err != nil {
		t.Fatalf("Не удалось создать таблицу пользователей: %v", err)
	}

	cleanup := func() {
		pool.Close()
		pgContainer.Terminate(ctx)
	}

	return pool, cleanup
}

func TestCreateUser(t *testing.T) {
	pool, cleanup := setupDBConfig(t)
	defer cleanup()

	repo := NewUserRepository(pool)

	newUser := &models.User{Name: "John Doe", Email: "john.doe@example.com"}
	err := repo.CreateUser(context.Background(), newUser)
	assert.NoError(t, err)
	assert.NotZero(t, newUser.ID)

	var fetchedUser models.User
	err = pool.QueryRow(context.Background(), "SELECT id, name, email FROM users WHERE id = $1", newUser.ID).
		Scan(&fetchedUser.ID, &fetchedUser.Name, &fetchedUser.Email)
	assert.NoError(t, err)
	assert.Equal(t, newUser.Name, fetchedUser.Name)
	assert.Equal(t, newUser.Email, fetchedUser.Email)

	duplicateUser := &models.User{Name: "Jane Doe", Email: "john.doe@example.com"}
	err = repo.CreateUser(context.Background(), duplicateUser)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "duplicate key value violates unique constraint")
}

func TestGetUserByID(t *testing.T) {
	pool, cleanup := setupDBConfig(t)
	defer cleanup()

	repo := NewUserRepository(pool)

	var userID int64
	err := pool.QueryRow(context.Background(), "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", "max", "max@example.com").Scan(&userID)
	assert.NoError(t, err)

	user, err := repo.GetUser(context.Background(), userID)
	assert.NoError(t, err)
	assert.Equal(t, userID, user.ID)
	assert.Equal(t, "max", user.Name)
	assert.Equal(t, "max@example.com", user.Email)

	_, err = repo.GetUser(context.Background(), 999)
	assert.Equal(t, errors.New("пользователь не найден"), err)
}

func TestUpdateUserByID(t *testing.T) {
	pool, cleanup := setupDBConfig(t)
	defer cleanup()

	repo := NewUserRepository(pool)

	var userID int64
	err := pool.QueryRow(context.Background(), "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", "Charlie", "charlie@example.com").Scan(&userID)
	assert.NoError(t, err)

	updateUser := &models.User{Name: "Charlie Updated", Email: "charlie.updated@example.com"}
	err = repo.UpdateUser(context.Background(), userID, updateUser)
	assert.NoError(t, err)

	var updatedUser models.User
	err = pool.QueryRow(context.Background(), "SELECT id, name, email FROM users WHERE id = $1", userID).
		Scan(&updatedUser.ID, &updatedUser.Name, &updatedUser.Email)
	assert.NoError(t, err)
	assert.Equal(t, "Charlie Updated", updatedUser.Name)
	assert.Equal(t, "charlie.updated@example.com", updatedUser.Email)

	err = repo.UpdateUser(context.Background(), 999, updateUser)
	assert.Equal(t, errors.New("пользователь не найден"), err)
}

func TestDeleteUserByID(t *testing.T) {
	pool, cleanup := setupDBConfig(t)
	defer cleanup()

	repo := NewUserRepository(pool)

	var userID int64
	err := pool.QueryRow(context.Background(), "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", "David", "david@example.com").Scan(&userID)
	assert.NoError(t, err)

	err = repo.DeleteUser(context.Background(), userID)
	assert.NoError(t, err)

	_, err = repo.GetUser(context.Background(), userID)
	assert.Equal(t, errors.New("пользователь не найден"), err)
}

func TestCreateUser_LongFields(t *testing.T) {
	pool, cleanup := setupDBConfig(t)
	defer cleanup()

	repo := NewUserRepository(pool)

	longName := strings.Repeat("a", 101)
	user := &models.User{Name: longName, Email: "longname@example.com"}
	err := repo.CreateUser(context.Background(), user)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "value too long for type character varying(100)")

	longEmail := strings.Repeat("b", 101) + "@example.com"
	user = &models.User{Name: "Long Email", Email: longEmail}
	err = repo.CreateUser(context.Background(), user)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "value too long for type character varying(100)")
}
