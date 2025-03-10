// internal/graphql/user/mutations.go
package userService

import (
	"schoolManagementSystem/api/graphql/services/userService/staff"
	"schoolManagementSystem/api/graphql/services/userService/student"

	"github.com/graphql-go/graphql"
)

// RootMutation returns all mutation fields related to users
func RootMutation() graphql.Fields {
	return graphql.Fields{
		"createStudent": student.CreateMutation,
		"adminLogin":    staff.LoginMutation,
	}
}
