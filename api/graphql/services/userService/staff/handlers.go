package staff

import (
	"context"
	"errors"
	"schoolManagementSystem/api/grpc"
	"schoolManagementSystem/pkg/helpers"
	pb "schoolManagementSystem/protos/auth"
	"time"

	"github.com/graphql-go/graphql"
)

var LoginHandler = func(p graphql.ResolveParams) (interface{}, error) {

	// parse TenantId from the request headers
	tenantId, err := helpers.GetTenantId(&p)
	if err != nil {
		return err, nil
	}

	uid, uidOk := p.Args["uid"].(string)
	password, passwordOk := p.Args["password"].(string)

	if !uidOk {
		return nil, errors.New("uid is required")
	}

	if !passwordOk {
		return nil, errors.New("password is required")
	}

	// Call to the Micro Service

	// Get gRPC client
	client, conn := grpc.InitAuthGRPC()
	defer conn.Close() // Ensure the connection is closed when done

	// Call the CreateUser gRPC method
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, errData := client.LoginStaff(ctx, &pb.LoginRequest{
		Uid:      uid,
		Password: password,
		TenantId: tenantId.String(),
	})

	if errData != nil {
		return nil, errData
	}

	// Return the created user data
	return res, nil
}
