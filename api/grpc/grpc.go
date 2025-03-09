package grpc

import (
	"flag"
	"log"

	pb "schoolManagementSystem/protos/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
)

// InitUserGRPC initializes the gRPC connection and returns a gRPC client and connection
func InitUserGRPC() (pb.UserServiceClient, *grpc.ClientConn) {
	// Connect to the gRPC server
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	// Initialize the client
	client := pb.NewUserServiceClient(conn)

	return client, conn
}
