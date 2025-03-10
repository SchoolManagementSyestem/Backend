package student

import (
	"schoolManagementSystem/api/graphql/common"

	"github.com/graphql-go/graphql"
)

// StudentType defines the GraphQL response for student creation
var StudentResponseType = common.GraphqlResponseWithData(
	"CreateStudent",
	graphql.Fields{
		"id":    &graphql.Field{Type: graphql.String},
		"email": &graphql.Field{Type: graphql.String},
	},
)
