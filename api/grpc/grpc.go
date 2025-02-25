package grpc

import (
	"context"
	"log"
	"time"

	pb "github.com/sumonskys/schoolManagementSystem/protos/user"
	"google.golang.org/grpc"
)

var UserServiceClient pb.UserServiceClient

// InitGRPC initializes the gRPC connection and returns a gRPC client
func InitGRPC() pb.UserServiceClient {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	// Defer closing the connection
	defer conn.Close()

	// Initialize the client
	UserServiceClient = pb.NewUserServiceClient(conn)

	return UserServiceClient
}

// CallGetUser is a helper function to call the GetUser RPC
func CallGetUser(userID string) (*pb.UserResponse, error) {
	// Ensure the client is initialized
	client := InitGRPC()

	// Call GetUser RPC method
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.UserRequest{UserId: userID}
	res, err := client.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
