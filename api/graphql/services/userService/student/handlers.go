package student

import (
	"context"
	"errors"
	"fmt"
	"schoolManagementSystem/api/grpc"
	"schoolManagementSystem/pkg/helpers"
	pb "schoolManagementSystem/protos/user"
	"time"

	"github.com/graphql-go/graphql"
)

var CreateStudentHandler = func(p graphql.ResolveParams) (interface{}, error) {

	// parse TenantId from the request headers
	tenantId, err := helpers.GetTenantId(&p)
	if err != nil {
		return err, nil
	}

	fmt.Println("Tenant ID: ", tenantId)

	firstName, firstNameOk := p.Args["firstName"].(string)
	email, _ := p.Args["email"].(string)

	if !firstNameOk {
		return nil, errors.New("firstName is required")
	}

	// Call to the Micro Service

	// Get gRPC client
	client, conn := grpc.InitUserGRPC()
	defer conn.Close() // Ensure the connection is closed when done

	// Call the CreateUser gRPC method
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, errData := client.CreateUser(ctx, &pb.CreateUserRequest{
		Name:  firstName,
		Email: email,
	})
	if errData != nil {
		return nil, errData
	}

	// Check if response and data exist
	if res == nil {
		return nil, errors.New("Failed to create user")
	}

	// Return the created user data
	return res, nil
}
