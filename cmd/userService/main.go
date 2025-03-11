package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"schoolManagementSystem/internal/db"
	pb "schoolManagementSystem/protos/user"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {

	// Load .env file
	envErr := godotenv.Load()
	if envErr != nil {
		log.Println("⚠️ No .env file found, using default system environment")
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT")))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Initialize dependencies
	database := db.InitPGDB()           // Database
	repo := NewUserRepository(database) // Database
	usecase := NewUserUsecase(repo)     // Business logic
	handler := NewUserHandler(usecase)  // gRPC Handler

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, handler)

	// Seed the database
	db.SeedUser(database)

	log.Println("User gRPC service running on port 50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
