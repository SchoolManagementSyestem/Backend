package staff

import (
	"context"
	"fmt"
	"schoolManagementSystem/api/graphql/common"
	"schoolManagementSystem/api/grpc"
	"schoolManagementSystem/pkg/helpers"
	pb "schoolManagementSystem/protos/auth"
	"time"

	"github.com/graphql-go/graphql"
)

var LoginAdminHandler = func(p graphql.ResolveParams) (interface{}, map[string]interface{}) {

	// parse TenantId from the request headers
	tenantId, err := helpers.GetTenantId(&p)
	if err != nil {
		return nil, err
	}

	fmt.Println("Tenant ID: ", tenantId)

	uid, uidOk := p.Args["uid"].(string)
	password, passwordOk := p.Args["password"].(string)

	if !uidOk {
		return nil, common.ResponseError("uid is required")
	}

	if !passwordOk {
		return nil, common.ResponseError("password is required")
	}

	// Call to the Micro Service

	// Get gRPC client
	client, conn := grpc.InitAuthGRPC()
	defer conn.Close() // Ensure the connection is closed when done

	// Call the CreateUser gRPC method
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, errData := client.LoginAdmin(ctx, &pb.LoginRequest{
		Uid:      uid,
		Password: password,
	})

	if errData != nil {
		return nil, common.ResponseError(fmt.Sprintf("failed to login: %v", err))
	}

	// Return the created user data
	return res, nil
}
