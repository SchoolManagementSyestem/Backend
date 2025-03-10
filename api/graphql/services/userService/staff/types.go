package staff

import (
	"schoolManagementSystem/api/graphql/common"

	"github.com/graphql-go/graphql"
)

// TokenType defines the GraphQL object for access and refresh tokens
var TokenType = graphql.NewObject(graphql.ObjectConfig{
	Name: "StaffLoginResponseToken",
	Fields: graphql.Fields{
		"accessToken":  &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"refreshToken": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
	},
})

// LoginResponseType defines the GraphQL response for staff login
var LoginAdminResponseType = common.GraphqlResponseWithData(
	"AdminLoginResponse",
	graphql.Fields{
		"id":             &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"firstName":      &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"lastName":       &graphql.Field{Type: graphql.String},
		"email":          &graphql.Field{Type: graphql.String},
		"phone":          &graphql.Field{Type: graphql.String},
		"dateOfBirth":    &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"gender":         &graphql.Field{Type: graphql.NewNonNull(common.GenderEnum)},
		"address":        &graphql.Field{Type: graphql.String},
		"status":         &graphql.Field{Type: graphql.NewNonNull(common.StatusEnum)},
		"profilePicture": &graphql.Field{Type: graphql.String},
		"createdAt":      &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"updatedAt":      &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"deletedAt":      &graphql.Field{Type: graphql.String},
		"role":           &graphql.Field{Type: graphql.NewNonNull(common.UserRoleEnum)},
		"token":          &graphql.Field{Type: graphql.NewNonNull(TokenType)},
	},
)
