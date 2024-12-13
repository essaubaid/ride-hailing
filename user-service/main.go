package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/essaubaid/ride-hailing/user-service/server"
)

func main() {
	grpcServer, listener := server.NewGRPCServer()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("Starting server on port 8090")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Could not start server %s", err)
		}
	}()

	<-stop

	log.Println("Shutting down server")
	grpcServer.GracefulStop()
	listener.Close()

	log.Println("Server stopped")
}
