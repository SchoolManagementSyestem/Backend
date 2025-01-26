// internal/graphql/schema.go
package graphql

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/sumonskys/schoolManagementSystem/internal/model"
	"github.com/sumonskys/schoolManagementSystem/internal/service"
)

type GraphQLHandler struct {
	userService *service.UserService
}

func NewGraphQLHandler(userService *service.UserService) *GraphQLHandler {
	return &GraphQLHandler{userService: userService}
}

// Define User type
var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":    &graphql.Field{Type: graphql.Int},
		"name":  &graphql.Field{Type: graphql.String},
		"email": &graphql.Field{Type: graphql.String},
	},
})

// Create RootQuery for fetching user
func (g *GraphQLHandler) RootQuery() graphql.Fields {
	return graphql.Fields{
		"getUser": &graphql.Field{
			Type:        userType,
			Description: "Get user by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)
				if !ok {
					return nil, errors.New("invalid ID format")
				}

				user, err := g.userService.GetUserByID(int64(id))
				if err != nil {
					return nil, err
				}
				return user, nil
			},
		},
	}
}

// Create RootMutation for creating user
func (g *GraphQLHandler) RootMutation() graphql.Fields {
	return graphql.Fields{
		"createUser": &graphql.Field{
			Type:        userType,
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

				user := &model.User{
					Name:  name,
					Email: email,
				}

				err := g.userService.CreateUser(user)
				if err != nil {
					return nil, err
				}
				return user, nil
			},
		},
	}
}

// Initialize and return the GraphQL schema
func (g *GraphQLHandler) InitSchema() (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    graphql.NewObject(graphql.ObjectConfig{Name: "Query", Fields: g.RootQuery()}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{Name: "Mutation", Fields: g.RootMutation()}),
	})
}
