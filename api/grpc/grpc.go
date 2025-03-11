package grpc

import (
	"log"
	"os"

	pbauth "schoolManagementSystem/protos/auth"
	pbuser "schoolManagementSystem/protos/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// InitUserGRPC initializes the gRPC connection and returns a gRPC client and connection
func InitUserGRPC() (pbuser.UserServiceClient, *grpc.ClientConn) {
	// Connect to the gRPC server
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(os.Getenv("USER_SERVICE_URL"), opts...)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	// Initialize the client
	client := pbuser.NewUserServiceClient(conn)

	return client, conn
}

// InitAuthGRPC initializes the gRPC connection and returns a gRPC client and connection
func InitAuthGRPC() (pbauth.AuthServiceClient, *grpc.ClientConn) {
	// Connect to the gRPC server
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(os.Getenv("AUTH_SERVICE_URL"), opts...)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	// Initialize the client
	client := pbauth.NewAuthServiceClient(conn)

	return client, conn
}
