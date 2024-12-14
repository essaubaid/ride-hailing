package main

import (
	"log"
	"os"

	"github.com/essaubaid/ride-hailing/common/db"
	"github.com/essaubaid/ride-hailing/common/server"
	"github.com/essaubaid/ride-hailing/proto/user"
	"github.com/essaubaid/ride-hailing/user-service/handlers"
	"github.com/essaubaid/ride-hailing/user-service/repositories"
	"github.com/essaubaid/ride-hailing/user-service/services"
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

	//Initialize the user repository
	userRepository := repositories.NewUserRepository(db)

	// Create a new gRPC server
	grpcConfig := server.GRPCServerConfig{Port: "8090"}
	grpcServer, listener := server.NewGRPCServer(&grpcConfig)

	handler := handlers.NewUserHandler(userRepository)
	UserService := services.NewUserService(*handler)

	// Register the user service with the gRPC server
	user.RegisterUserServiceServer(grpcServer, UserService)

	// Run the server with graceful shutdown.
	server.RunGRPCServer(grpcServer, listener, "UserService")
}
