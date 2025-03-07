// internal/graphql/user/types.go
package student

import "github.com/graphql-go/graphql"

// StudentType defines the GraphQL object for User
var StudentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":    &graphql.Field{Type: graphql.Int},
		"name":  &graphql.Field{Type: graphql.String},
		"email": &graphql.Field{Type: graphql.String},
	},
})
