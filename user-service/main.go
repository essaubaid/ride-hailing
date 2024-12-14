package main

import (
	"github.com/essaubaid/ride-hailing/common/server"
	"github.com/essaubaid/ride-hailing/proto/user"
	"github.com/essaubaid/ride-hailing/user-service/handlers"
	"github.com/essaubaid/ride-hailing/user-service/services"
)

func main() {
	grpcConfig := server.GRPCServerConfig{Port: "8090"}
	grpcServer, listener := server.NewGRPCServer(&grpcConfig)

	handler := handlers.NewUserHandler()
	UserService := services.NewUserService(*handler)

	// Register the user service with the gRPC server
	user.RegisterUserServiceServer(grpcServer, UserService)

	// Run the server with graceful shutdown.
	server.RunGRPCServer(grpcServer, listener, "UserService")
}
