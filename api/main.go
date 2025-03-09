package main

import (
	"log"
	"net/http"

	"schoolManagementSystem/api/graphql"
	"schoolManagementSystem/api/middleware"
	"schoolManagementSystem/api/rest"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize GraphQL handler
	graphqlHandler := graphql.NewGraphQLHandler()

	// Create HTTP router
	router := mux.NewRouter()
	router.Handle("/graphql", middleware.AuthMiddleware(graphqlHandler))

	// Register REST API routes
	rest.RegisterRoutes(router)

	// Start server
	log.Println("Server running at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
