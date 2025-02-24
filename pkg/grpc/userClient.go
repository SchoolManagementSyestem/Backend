package grpc

import (
	"context"
	"log"

	pb "github.com/sumonskys/schoolManagementSystem/api/proto/user"
	"google.golang.org/grpc"
)

// UserClient wraps the gRPC client connection
type UserClient struct {
	client pb.UserServiceClient
}

// NewUserClient initializes a gRPC client
func NewUserClient(grpcAddress string) *UserClient {
	conn, err := grpc.Dial(grpcAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to User service: %v", err)
	}
	return &UserClient{
		client: pb.NewUserServiceClient(conn),
	}
}

// GetUser calls the gRPC method to get a user
func (uc *UserClient) GetUser(userId string) (*pb.UserResponse, error) {
	req := &pb.UserRequest{UserId: userId}
	return uc.client.GetUser(context.Background(), req)
}
