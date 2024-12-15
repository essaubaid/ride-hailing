package handlers

import (
	"context"

	"github.com/essaubaid/ride-hailing/common/models"
	"github.com/essaubaid/ride-hailing/user-service/repositories"
)

type UserHandlerImpl struct {
	repo *repositories.UserRepository
}

func NewUserHandlerImpl(repo *repositories.UserRepository) *UserHandlerImpl {
	return &UserHandlerImpl{
		repo: repo,
	}
}

func (h *UserHandlerImpl) GetUser(ctx context.Context, userId int32) (*models.User, error) {

	user, err := h.repo.GetUserByID(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (h *UserHandlerImpl) CreateUser(ctx context.Context, name string) (*models.User, error) {

	user := models.User{Name: name}
	_, err := h.repo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (h *UserHandlerImpl) DeleteUser(ctx context.Context, userId int32) error {

	if err := h.repo.DeleteUser(userId); err != nil {
		return err
	}

	return nil
}
