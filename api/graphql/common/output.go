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

// ResponseWithData creates a reusable response with a dynamic data field.
func ResponseWithData(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  true,
		"message": message,
		"data":    data,
	}
}

// ResponseSuccess creates a reusable success response.
func ResponseSuccess(message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  true,
		"message": message,
	}
}

// ResponseError creates a reusable error response.
func ResponseError(message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  false,
		"message": message,
	}
}
