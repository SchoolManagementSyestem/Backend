package grpc

import (
	"context"
	"log"

	pb "github.com/sumonskys/schoolManagementSystem/api/proto"
	"google.golang.org/grpc"
)

// AuthClient wraps the gRPC client connection
type AuthClient struct {
	client pb.AuthServiceClient
}

// NewAuthClient initializes a gRPC client
func NewAuthClient(grpcAddress string) *AuthClient {
	conn, err := grpc.Dial(grpcAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Auth service: %v", err)
	}
	return &AuthClient{
		client: pb.NewAuthServiceClient(conn),
	}
}

// Authenticate calls gRPC method
func (ac *AuthClient) Authenticate(username, password string) (*pb.AuthResponse, error) {
	req := &pb.AuthRequest{
		Username: username,
		Password: password,
	}
	return ac.client.Authenticate(context.Background(), req)
}
