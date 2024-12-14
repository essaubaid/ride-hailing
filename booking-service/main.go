package main

import (
	"github.com/essaubaid/ride-hailing/booking-service/handlers"
	"github.com/essaubaid/ride-hailing/booking-service/services"
	"github.com/essaubaid/ride-hailing/common/server"
	"github.com/essaubaid/ride-hailing/proto/booking"
	"github.com/essaubaid/ride-hailing/proto/rides"
)

func main() {
	grpcConfig := server.GRPCServerConfig{Port: "8092"}
	grpcServer, listener := server.NewGRPCServer(&grpcConfig)

	ridesHandler := handlers.NewRidesHandler()
	ridesService := services.NewRidesService(*ridesHandler)
	bookingHandler := handlers.NewBookingHandler()
	bookingService := services.NewBookingService(*bookingHandler)

	// Register the rides service with the gRPC server
	rides.RegisterRidesServiceServer(grpcServer, ridesService)
	// Register the booking service with the gRPC server
	booking.RegisterBookingServiceServer(grpcServer, bookingService)

	// Run the server with graceful shutdown.
	server.RunGRPCServer(grpcServer, listener, "BookingService")
}
