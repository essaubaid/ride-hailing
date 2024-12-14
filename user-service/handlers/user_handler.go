package handlers

import (
	"context"

	"github.com/essaubaid/ride-hailing/common/models"
	"github.com/essaubaid/ride-hailing/user-service/repositories"
)

type UserHandler struct {
	repo *repositories.UserRepository
}

func NewUserHandler(repo *repositories.UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

func (h *UserHandler) GetUser(ctx context.Context, userId int32) (*models.User, error) {

	user, err := h.repo.GetUserByID(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (h *UserHandler) CreateUser(ctx context.Context, name string) (*models.User, error) {

	user := models.User{Name: name}
	_, err := h.repo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, userId int32) error {

	if err := h.repo.DeleteUser(userId); err != nil {
		return err
	}

	return nil
}
