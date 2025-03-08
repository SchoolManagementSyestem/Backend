package student

import (
	"context"
	"errors"
	"fmt"
	"schoolManagementSystem/api/grpc"
	"schoolManagementSystem/helpers"
	pb "schoolManagementSystem/protos/user"
	"time"

	"github.com/graphql-go/graphql"
)

var CreateStudentHandler = func(p graphql.ResolveParams) (interface{}, error) {

	// parse TenantId from the request headers
	tenantId, err := helpers.GetTenantId(&p)
	if err != nil {
		return nil, err
	}

	// Check The Permission

	fmt.Println(tenantId)
	name, nameOk := p.Args["name"].(string)
	email, emailOk := p.Args["email"].(string)

	if !nameOk || !emailOk {
		return nil, errors.New("invalid input")
	}

	// Call to the Micro Service

	// Get gRPC client
	client, conn := grpc.InitUserGRPC()
	defer conn.Close() // Ensure the connection is closed when done

	// Call the CreateUser gRPC method
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := client.CreateUser(ctx, &pb.CreateUserRequest{
		Name:  name,
		Email: email,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	// Check if response and data exist
	if res == nil || res.Data == nil {
		return nil, errors.New("invalid response from gRPC server")
	}

	// Return the created user data
	return map[string]interface{}{
		"id":    res.GetData().GetId(),
		"name":  res.GetData().GetName(),
		"email": res.GetData().GetEmail(),
	}, nil
}
