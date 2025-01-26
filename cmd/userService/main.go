// cmd/userService/main.go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/graphql-go/handler"
	"github.com/segmentio/kafka-go"
	"github.com/sumonskys/schoolManagementSystem/internal/graphql"
	"github.com/sumonskys/schoolManagementSystem/internal/repository"
	"github.com/sumonskys/schoolManagementSystem/internal/service"
)

func main() {
	producer := initKafkaProducer()
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo, producer)
	graphqlHandler := graphql.NewGraphQLHandler(userService)

	// Create GraphQL schema
	schema, err := graphqlHandler.InitSchema()
	if err != nil {
		log.Fatalf("failed to initialize schema: %v", err)
	}

	// GraphQL handler
	graphqlHandlerFunc := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// Create router
	router := mux.NewRouter()
	router.Handle("/graphql", graphqlHandlerFunc)

	// Start server
	log.Println("Server running at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func initKafkaProducer() *kafka.Writer {
	// Initialize Kafka producer
	return &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "user_topic",
		Balancer: &kafka.LeastBytes{},
	}
}
