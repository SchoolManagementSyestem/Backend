package auth

import (
	"encoding/json"
	"net/http"
)

// AuthController handles authentication-related HTTP requests
type AuthController struct {
	authService *AuthService
}

// NewAuthController initializes the AuthController
func NewAuthController(authService *AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// LoginHandler is the endpoint for logging in
func (ac *AuthController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Parse the request body
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call AuthService to authenticate user
	token, err := ac.authService.Authenticate(credentials.Username, credentials.Password)
	if err != nil {
		http.Error(w, "Authentication failed", http.StatusUnauthorized)
		return
	}

	// Respond with the JWT token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
