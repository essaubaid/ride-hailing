package server

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/essaubaid/ride-hailing/common/logging"
	"google.golang.org/grpc"
)

var logger = logging.GetLogger()

type GRPCServerConfig struct {
	Port string
}

func NewGRPCServer(config *GRPCServerConfig) (*grpc.Server, net.Listener) {
	// Create a new listener on the specified port
	listener, err := net.Listen("tcp", ":"+config.Port)
	if err != nil {
		logger.Fatalf("Failed to start server: %v", err)
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
		logger.Infof("Starting %s on port %s", serviceName, listener.Addr().String())
		if err := grpcServer.Serve(listener); err != nil {
			logger.Fatalf("Failed to serve %s: %v", serviceName, err)
		}
	}()

	// Wait for shutdown signal.
	<-stop
	logger.Infof("Stopping %s...", serviceName)
	grpcServer.GracefulStop()
	logger.Infof("%s stopped", serviceName)
}
