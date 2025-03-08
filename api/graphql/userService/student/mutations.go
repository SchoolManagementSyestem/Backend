package student

import (
	"github.com/graphql-go/graphql"
)

// CreateStudentMutation resolver
var CreateStudentMutation = &graphql.Field{
	Type:        StudentResponseType,
	Description: "Create a new student",
	Args:        CreateStudentArgs,
	Resolve:     CreateStudentHandler,
}
