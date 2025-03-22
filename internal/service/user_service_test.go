package service

import (
	"context"
	"errors"
	"testing"
	"testovoe/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) UpdateUser(ctx context.Context, id int64, user *models.User) error {
	args := m.Called(ctx, id, user)
	return args.Error(0)
}

func (m *MockUserRepository) DeleteUser(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
func (m *MockUserRepository) GetUser(ctx context.Context, id int64) (*models.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.User), args.Error(1)
}


func TestCreateUser_EmptyFields(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	user := &models.User{Name: "", Email: "test@example.com"}
	err := service.CreateUser(context.Background(), user)
	assert.Equal(t, errors.New("Имя или email не могут быть пустыми"), err)

	user = &models.User{Name: "Test User", Email: ""}
	err = service.CreateUser(context.Background(), user)
	assert.Equal(t, errors.New("Имя или email не могут быть пустыми"), err)

	mockRepo.AssertNotCalled(t, "CreateUser")
}
func TestCreateUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	user := &models.User{Name: "Test User", Email: "test@example.com"}
	mockRepo.On("CreateUser", mock.Anything, user).Return(nil)

	err := service.CreateUser(context.Background(), user)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetUserByID_NotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	mockRepo.On("GetUserByID", mock.Anything, int64(1)).Return((*models.User)(nil), errors.New("пользователь не найден"))

	user, err := service.GetUser(context.Background(), 1)
	assert.Equal(t, errors.New("пользователь не найден"), err)
	assert.Nil(t, user)
	mockRepo.AssertExpectations(t)
}
func TestGetUserByID_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	expectedUser := &models.User{ID: 1, Name: "Test User", Email: "test@example.com"}
	mockRepo.On("GetUserByID", mock.Anything, int64(1)).Return(expectedUser, nil)

	user, err := service.GetUser(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUserByID_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	user := &models.User{Name: "Updated User", Email: "updated@example.com"}
	mockRepo.On("UpdateUserByID", mock.Anything, int64(1), user).Return(nil)

	err := service.UpdateUser(context.Background(), 1, user)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUserByID_EmptyFields(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	user := &models.User{Name: "", Email: "updated@example.com"}
	err := service.UpdateUser(context.Background(), 1, user)
	assert.Equal(t, errors.New("Имя или email не могут быть пустыми"), err)

	user = &models.User{Name: "Updated User", Email: ""}
	err = service.UpdateUser(context.Background(), 1, user)
	assert.Equal(t, errors.New("Имя или email не могут быть пустыми"), err)

	mockRepo.AssertNotCalled(t, "UpdateUserByID")
}

func TestDeleteUserByID_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	mockRepo.On("DeleteUserByID", mock.Anything, int64(1)).Return(nil)

	err := service.DeleteUser(context.Background(), 1)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteUserByID_NotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	mockRepo.On("DeleteUserByID", mock.Anything, int64(1)).Return(errors.New("пользователь не найден"))

	err := service.DeleteUser(context.Background(), 1)
	assert.Equal(t,errors.New("пользователь не найден"), err)
	mockRepo.AssertExpectations(t)
}
