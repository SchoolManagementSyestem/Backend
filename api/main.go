package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sumonskys/schoolManagementSystem/api/graphql"
	"github.com/sumonskys/schoolManagementSystem/api/rest"
)

func main() {
	// Initialize GraphQL handler
	graphqlHandler := graphql.NewGraphQLHandler()

	// Create HTTP router
	router := mux.NewRouter()
	router.Handle("/graphql", graphqlHandler) // Attach the handler

	// Register REST API routes
	rest.RegisterRoutes(router)

	// Start server
	log.Println("Server running at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
