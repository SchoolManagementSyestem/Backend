package rest

import (
	"schoolManagementSystem/internal/auth"
	"schoolManagementSystem/internal/user"

	"github.com/gorilla/mux"
)

// RegisterRoutes sets up REST API routes
func RegisterRoutes(router *mux.Router) {
	// Initialize controllers with gRPC service address and secret key
	authController := auth.NewAuthController(auth.NewAuthService("mysecretkey"))
	userController := user.NewUserController()

	// Define routes
	router.HandleFunc("/api/auth/login", authController.LoginHandler).Methods("POST")
	router.HandleFunc("/api/user/{userId}", userController.GetUserHandler).Methods("GET")
}
