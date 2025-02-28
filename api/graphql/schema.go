// internal/graphql/schema.go
package graphql

import (
	"schoolManagementSystem/api/graphql/user"

	"github.com/graphql-go/graphql"
)

// InitSchema initializes the GraphQL schema
func InitSchema() (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: mergeFields(
				user.RootQuery(),
			),
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: mergeFields(
				user.RootMutation(),
			),
		}),
	})
}

// mergeFields combines multiple GraphQL fields
func mergeFields(fields ...graphql.Fields) graphql.Fields {
	merged := graphql.Fields{}
	for _, field := range fields {
		for key, value := range field {
			merged[key] = value
		}
	}
	return merged
}
