package student

import (
	"schoolManagementSystem/api/graphql/common"

	"github.com/graphql-go/graphql"
)

var CreateStudentArgs = graphql.FieldConfigArgument{
	"firstName":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	"lastName":       &graphql.ArgumentConfig{Type: graphql.String},
	"email":          &graphql.ArgumentConfig{Type: graphql.String},
	"phone":          &graphql.ArgumentConfig{Type: graphql.String},
	"dateOfBirth":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	"gender":         &graphql.ArgumentConfig{Type: graphql.NewNonNull(common.GenderEnum)},
	"address":        &graphql.ArgumentConfig{Type: graphql.String},
	"profilePicture": &graphql.ArgumentConfig{Type: graphql.String},
	"status":         &graphql.ArgumentConfig{Type: graphql.NewNonNull(common.StatusEnum)},
}
