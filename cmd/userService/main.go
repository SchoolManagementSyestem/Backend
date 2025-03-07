package main

import (
	"log"
	"net"

	"schoolManagementSystem/internal/db"
	pb "schoolManagementSystem/protos/user"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Load .env file
	envErr := godotenv.Load()
	if envErr != nil {
		log.Println("⚠️ No .env file found, using default system environment")
	}

	// Initialize dependencies
	database := db.InitPGDB()           // Database
	repo := NewUserRepository(database) // Database
	usecase := NewUserUsecase(repo)     // Business logic
	handler := NewUserHandler(usecase)  // gRPC Handler

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, handler)

	// Seed the database
	db.Seed(database)

	log.Println("User gRPC service running on port 50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
