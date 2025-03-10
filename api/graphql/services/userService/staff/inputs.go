package staff

import (
	"github.com/graphql-go/graphql"
)

var LoginAdminArgs = graphql.FieldConfigArgument{
	"uid":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	"status": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
}
