package staff

import (
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"
)

var GetStudentQuery = &graphql.Field{
	Type:        LoginAdminResponseType,
	Description: "Get user by ID",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{Type: graphql.Int},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, ok := p.Args["id"].(int)
		if !ok {
			return nil, errors.New("invalid ID format")
		}

		// Simulated user lookup
		fmt.Println("Fetching user with ID:", id)
		return nil, nil
	},
}
