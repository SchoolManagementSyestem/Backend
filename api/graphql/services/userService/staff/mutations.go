package staff

import (
	"github.com/graphql-go/graphql"
)

// LoginMutation resolver
var LoginMutation = &graphql.Field{
	Type:        LoginResponseType,
	Description: "Staff login",
	Args:        LoginArgs,
	Resolve:     LoginHandler,
}
