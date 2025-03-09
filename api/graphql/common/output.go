package common

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

// GraphqlResponseWithData creates a reusable response type with a dynamic data field.
func GraphqlResponseWithData(name string, data graphql.Fields) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: fmt.Sprintf("%sResponse", name),
		Fields: graphql.Fields{
			"status":  &graphql.Field{Type: graphql.Boolean},
			"message": &graphql.Field{Type: graphql.String},
			"data": &graphql.Field{Type: graphql.NewObject(graphql.ObjectConfig{
				Name:   fmt.Sprintf("%sResponseData", name),
				Fields: data,
			})},
		},
	})
}

// GraphqlResponse creates a reusable response.
func GraphqlResponse(name string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: fmt.Sprintf("%sResponse", name),
		Fields: graphql.Fields{
			"status":  &graphql.Field{Type: graphql.Boolean},
			"message": &graphql.Field{Type: graphql.String},
		},
	})
}
