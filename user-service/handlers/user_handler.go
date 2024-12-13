package handlers

import "context"

type User struct {
	Id   int32
	Name string
}

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetUser(ctx context.Context, userId int32) (*User, error) {

	return &User{
		Id:   userId,
		Name: "John Doe",
	}, nil
}

func (h *UserHandler) CreateUser(ctx context.Context, name string) (*User, error) {

	return &User{
		Id:   1,
		Name: name,
	}, nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, userId int32) error {

	return nil
}
