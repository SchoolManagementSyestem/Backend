package main

import (
	"log"
	"net"

	pb "schoolManagementSystem/protos/user"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Initialize dependencies
	repo := NewUserRepository()        // Database
	usecase := NewUserUsecase(repo)    // Business logic
	handler := NewUserHandler(usecase) // gRPC Handler

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, handler)

	log.Println("User gRPC service running on port 50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
