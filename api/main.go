package main

import (
	"log"
	"net/http"

	"schoolManagementSystem/api/graphql"
	"schoolManagementSystem/api/middleware"
	"schoolManagementSystem/api/rest"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	envErr := godotenv.Load("../.env")
	if envErr != nil {
		log.Fatal("⚠️ No .env file found, using default system environment")
	}
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
