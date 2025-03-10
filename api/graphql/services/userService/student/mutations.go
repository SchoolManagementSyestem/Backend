package student

import (
	"github.com/graphql-go/graphql"
)

// CreateMutation resolver
var CreateMutation = &graphql.Field{
	Type:        StudentResponseType,
	Description: "Create a new student",
	Args:        CreateStudentArgs,
	Resolve:     CreateStudentHandler,
}
