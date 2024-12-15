package services

import (
	"context"
	"testing"

	"github.com/essaubaid/ride-hailing/common/models"
	"github.com/essaubaid/ride-hailing/proto/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock for UserHandler
type MockUserHandler struct {
	mock.Mock
}

func (m *MockUserHandler) GetUser(ctx context.Context, userId int32) (*models.User, error) {
	args := m.Called(ctx, userId)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserHandler) CreateUser(ctx context.Context, name string) (*models.User, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserHandler) DeleteUser(ctx context.Context, userId int32) error {
	args := m.Called(ctx, userId)
	return args.Error(0)
}

func TestUserService_GetUser(t *testing.T) {
	// Create a mock handler
	mockHandler := new(MockUserHandler)

	// Create a test instance of UserService
	service := NewUserService(mockHandler)

	// Define test inputs and expected outputs
	userId := int32(1)
	expectedUser := &models.User{Name: "John Doe"}

	// Set up mock expectations
	mockHandler.On("GetUser", mock.Anything, userId).Return(expectedUser, nil)

	// Call the service method
	req := &user.GetUserRequest{Id: userId}
	resp, err := service.GetUser(context.Background(), req)

	// Assert the results
	assert.NoError(t, err)
	assert.Equal(t, expectedUser.Name, resp.Name)
	mockHandler.AssertExpectations(t)
}

func TestUserService_CreateUser(t *testing.T) {
	// Create a mock handler
	mockHandler := new(MockUserHandler)

	// Create a test instance of UserService
	service := NewUserService(mockHandler)

	// Define test inputs and expected outputs
	userName := "John Doe"
	expectedUser := &models.User{Id: 1, Name: userName}

	// Set up mock expectations
	mockHandler.On("CreateUser", mock.Anything, userName).Return(expectedUser, nil)

	// Call the service method
	req := &user.CreateUserRequest{Name: userName}
	resp, err := service.CreateUser(context.Background(), req)

	// Assert the results
	assert.NoError(t, err)
	assert.Equal(t, expectedUser.Id, resp.Id)
	mockHandler.AssertExpectations(t)
}

func TestUserService_DeleteUser(t *testing.T) {
	// Create a mock handler
	mockHandler := new(MockUserHandler)

	// Create a test instance of UserService
	service := NewUserService(mockHandler)

	// Define test inputs
	userId := int32(1)

	// Set up mock expectations
	mockHandler.On("DeleteUser", mock.Anything, userId).Return(nil)

	// Call the service method
	req := &user.DeleteUserRequest{Id: userId}
	resp, err := service.DeleteUser(context.Background(), req)

	// Assert the results
	assert.NoError(t, err)
	assert.Equal(t, "User deleted successfully", resp.Message)
	mockHandler.AssertExpectations(t)
}
