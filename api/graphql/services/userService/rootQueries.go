// internal/graphql/user/queries.go
package userService

import (
	"schoolManagementSystem/api/graphql/services/userService/student"

	"github.com/graphql-go/graphql"
)

// RootQuery returns all query fields related to users
func RootQuery() graphql.Fields {
	return graphql.Fields{
		"getStudent": student.GetStudentQuery,
	}
}
