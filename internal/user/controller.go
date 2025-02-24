package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// UserController handles HTTP requests related to users
type UserController struct {
	userService *UserService
}

// NewUserController initializes the UserController
func NewUserController() *UserController {
	return &UserController{
		userService: NewUserService(), // Initialize UserService
	}
}

// GetUserHandler handles fetching user details by ID
func (uc *UserController) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	// Fetch user details from UserService
	user, err := uc.userService.GetUser(userId)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Respond with user details in JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
