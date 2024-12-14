package main

import (
	"log"
	"os"

	"github.com/essaubaid/ride-hailing/booking-service/handlers"
	"github.com/essaubaid/ride-hailing/booking-service/repositories"
	"github.com/essaubaid/ride-hailing/booking-service/services"
	"github.com/essaubaid/ride-hailing/common/db"
	"github.com/essaubaid/ride-hailing/common/server"
	"github.com/essaubaid/ride-hailing/proto/booking"
	"github.com/essaubaid/ride-hailing/proto/rides"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Create DB connection
	dbConfig := db.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     5432,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
	}

	db, err := db.NewDatabase(dbConfig)
	if err != nil {
		log.Fatalf("could not connect to DB: %v", err)
	}
	defer db.Close()

	//Initialize the repositories
	ridesRepository := repositories.NewRidesRepository(db)

	// Create a new gRPC server
	grpcConfig := server.GRPCServerConfig{Port: "8092"}
	grpcServer, listener := server.NewGRPCServer(&grpcConfig)

	ridesHandler := handlers.NewRidesHandler(ridesRepository)
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
