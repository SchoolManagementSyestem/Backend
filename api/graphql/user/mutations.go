// internal/graphql/user/mutations.go
package user

import (
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"
)

// CreateUserMutation resolver
var CreateUserMutation = &graphql.Field{
	Type:        UserType,
	Description: "Create a new user",
	Args: graphql.FieldConfigArgument{
		"name":  &graphql.ArgumentConfig{Type: graphql.String},
		"email": &graphql.ArgumentConfig{Type: graphql.String},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		name, nameOk := p.Args["name"].(string)
		email, emailOk := p.Args["email"].(string)

		if !nameOk || !emailOk {
			return nil, errors.New("invalid input")
		}

		// Simulated user creation
		fmt.Println("Creating user:", name, email)
		return nil, nil
	},
}

// RootMutation returns all mutation fields related to users
func RootMutation() graphql.Fields {
	return graphql.Fields{
		"createUser": CreateUserMutation,
	}
}
