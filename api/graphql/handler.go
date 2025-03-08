// internal/graphql/handlers.go
package graphql

import (
	"context"
	"log"
	"net/http"

	"github.com/graphql-go/handler"
)

// GraphQLHandler struct
type GraphQLHandler struct {
	handler *handler.Handler
}

// NewGraphQLHandler initializes the GraphQL handler
func NewGraphQLHandler() *GraphQLHandler {
	schema, err := InitSchema()
	if err != nil {
		log.Fatalf("Failed to initialize GraphQL schema: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true, // Enable GraphiQL UI for easy testing
	})

	return &GraphQLHandler{handler: h}
}

// ServeHTTP makes GraphQLHandler compatible with http.Handler
func (g *GraphQLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Create a new context with the HTTP request attached
	ctx := context.WithValue(r.Context(), "httpRequest", r)

	// Call the handler with the new context
	r = r.WithContext(ctx)

	// Now, process the request with the updated context
	g.handler.ServeHTTP(w, r)
}
