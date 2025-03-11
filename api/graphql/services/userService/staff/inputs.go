package staff

import (
	"github.com/graphql-go/graphql"
)

var LoginArgs = graphql.FieldConfigArgument{
	"uid":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
}
