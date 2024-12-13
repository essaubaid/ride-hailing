package main

import (
	"context"
	"log"
	"net"

	"github.com/essaubaid/ride-hailing/proto/user"
	"google.golang.org/grpc"
)

type myUserServiceServer struct {
	user.UnimplementedUserServiceServer
}

func (server myUserServiceServer) GetUser(ctx context.Context, req *user.UserRequest) (*user.UserResponse, error) {
	log.Printf("Received request for user ID: %v", req.Id)

	return &user.UserResponse{Name: "John Doe", Email: "john@example.com"}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8091")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &myUserServiceServer{}

	user.RegisterUserServiceServer(serverRegistrar, service)

	if err := serverRegistrar.Serve(listener); err != nil {
		log.Fatalf("Could not start server %s", err)
	}
}
