package server

import (
	"log"
	"net"

	"github.com/essaubaid/ride-hailing/proto/user"
	"github.com/essaubaid/ride-hailing/user-service/handlers"
	"github.com/essaubaid/ride-hailing/user-service/services"
	"google.golang.org/grpc"
)

func NewGRPCServer() (*grpc.Server, net.Listener) {
	// Create a new listener on port 8090
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	handler := handlers.NewUserHandler()
	UserService := services.NewUserService(*handler)

	// Register the user service with the gRPC server
	user.RegisterUserServiceServer(grpcServer, UserService)

	return grpcServer, listener
}
