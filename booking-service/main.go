package main

import (
	"os"

	"github.com/essaubaid/ride-hailing/booking-service/handlers"
	"github.com/essaubaid/ride-hailing/booking-service/repositories"
	"github.com/essaubaid/ride-hailing/booking-service/services"
	"github.com/essaubaid/ride-hailing/common/db"
	"github.com/essaubaid/ride-hailing/common/logging"
	"github.com/essaubaid/ride-hailing/common/server"
	"github.com/essaubaid/ride-hailing/proto/booking"
	"github.com/essaubaid/ride-hailing/proto/rides"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var logger = logging.GetLogger()

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		logger.Fatalf("Error loading .env file")
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
		logger.Fatalf("could not connect to DB: %v", err)
	}
	defer db.Close()

	// Connect to RideService.
	rideConn, err := grpc.NewClient("ride_service:80", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatalf("Failed to connect to RideService: %v", err)
	}
	defer rideConn.Close()
	rideClient := rides.NewRidesServiceClient(rideConn)

	//Initialize the repositories
	bookingRepository := repositories.NewBookingRepository(db)

	// Create a new gRPC server
	grpcConfig := server.GRPCServerConfig{Port: os.Getenv("SERVICE_PORT")}
	grpcServer, listener := server.NewGRPCServer(&grpcConfig)

	bookingHandler := handlers.NewBookingHandler(bookingRepository, &rideClient)
	bookingService := services.NewBookingService(*bookingHandler)

	// Register the booking service with the gRPC server
	booking.RegisterBookingServiceServer(grpcServer, bookingService)

	// Run the server with graceful shutdown.
	server.RunGRPCServer(grpcServer, listener, "BookingService")
}
