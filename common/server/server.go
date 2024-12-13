package server

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

type GRPCServerConfig struct {
	Port string
}

func NewGRPCServer(config *GRPCServerConfig) (*grpc.Server, net.Listener) {
	// Create a new listener on the specified port
	listener, err := net.Listen("tcp", ":"+config.Port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	return grpcServer, listener
}

// RunGRPCServer starts the gRPC server and handles graceful shutdown.
func RunGRPCServer(grpcServer *grpc.Server, listener net.Listener, serviceName string) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Starting %s on port %s", serviceName, listener.Addr().String())
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve %s: %v", serviceName, err)
		}
	}()

	// Wait for shutdown signal.
	<-stop
	log.Printf("Stopping %s...", serviceName)
	grpcServer.GracefulStop()
	log.Printf("%s stopped", serviceName)
}
