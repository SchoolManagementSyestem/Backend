package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"schoolManagementSystem/internal/db"
	pb "schoolManagementSystem/protos/auth"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {

	// Load .env file
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("⚠️ No .env file found, using default system environment")
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT")))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Initialize dependencies
	database := db.InitPGDB()           // Database
	repo := NewAuthRepository(database) // Database
	usecase := NewAuthUsecase(repo)     // Business logic
	handler := NewAuthHandler(usecase)  // gRPC Handler

	server := grpc.NewServer()
	pb.RegisterAuthServiceServer(server, handler)

	log.Println("User gRPC service running on port 50052")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
