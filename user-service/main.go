package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/essaubaid/ride-hailing/proto/user"
	"google.golang.org/grpc"
)

type myUserServiceServer struct {
	user.UnimplementedUserServiceServer
}

func (server myUserServiceServer) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	log.Printf("Received request for user ID: %v", req.Id)

	return &user.GetUserResponse{Name: "John Doe"}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &myUserServiceServer{}

	user.RegisterUserServiceServer(serverRegistrar, service)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("Starting server on port 8090")
		if err := serverRegistrar.Serve(listener); err != nil {
			log.Fatalf("Could not start server %s", err)
		}
	}()

	<-stop

	log.Println("Shutting down server")
	serverRegistrar.GracefulStop()
	listener.Close()

	log.Println("Server stopped")
}
