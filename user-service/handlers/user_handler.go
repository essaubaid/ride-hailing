package handlers

import (
	"context"

	"github.com/essaubaid/ride-hailing/common/models"
)

type UserHandler interface {
	GetUser(ctx context.Context, userId int32) (*models.User, error)
	CreateUser(ctx context.Context, name string) (*models.User, error)
	DeleteUser(ctx context.Context, userId int32) error
}
