package services

import (
	"context"
	"log"

	"github.com/essaubaid/ride-hailing/proto/user"
	"github.com/essaubaid/ride-hailing/user-service/handlers"
)

type UserService struct {
	user.UnimplementedUserServiceServer
	handler handlers.UserHandler
}

func NewUserService(handler handlers.UserHandler) *UserService {
	return &UserService{
		handler: handler,
	}
}

func (s *UserService) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	log.Printf("gRPC: Received GetUser request for ID: %d", req.Id)

	// Call the handler to get the user
	userData, err := s.handler.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &user.GetUserResponse{
		Name: userData.Name,
	}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	log.Printf("gRPC: Received CreateUser request for name: %s", req.Name)

	// Call the handler to create the user
	userData, err := s.handler.CreateUser(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &user.CreateUserResponse{
		Id: userData.Id,
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	log.Printf("gRPC: Received DeleteUser request for ID: %d", req.Id)

	// Call the handler to delete the user
	err := s.handler.DeleteUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &user.DeleteUserResponse{
		Message: "User deleted successfully",
	}, nil
}
