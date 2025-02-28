// internal/graphql/user/mutations.go
package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"schoolManagementSystem/api/grpc"
	pb "schoolManagementSystem/protos/user"

	"github.com/graphql-go/graphql"
)

// CreateUserMutation resolver
var CreateUserMutation = &graphql.Field{
	Type:        UserType,
	Description: "Create a new user",
	Args: graphql.FieldConfigArgument{
		"name":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"email": &graphql.ArgumentConfig{Type: graphql.String},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
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
	},
}

// RootMutation returns all mutation fields related to users
func RootMutation() graphql.Fields {
	return graphql.Fields{
		"createUser": CreateUserMutation,
	}
}
