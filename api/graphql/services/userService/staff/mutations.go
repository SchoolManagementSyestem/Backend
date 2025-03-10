package staff

import (
	"github.com/graphql-go/graphql"
)

// LoginMutation resolver
var LoginMutation = &graphql.Field{
	Type:        LoginAdminResponseType,
	Description: "Admin login",
	Args:        LoginAdminArgs,
	Resolve:     LoginAdminHandler,
}
