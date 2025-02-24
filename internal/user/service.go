package user

import (
	"errors"
)

// UserService handles the user-related business logic
type UserService struct {
	users map[string]User
}

// User defines a user entity
type User struct {
	Username string
	Email    string
}

// NewUserService initializes the user service
func NewUserService() *UserService {
	return &UserService{
		users: map[string]User{
			"123": {Username: "admin", Email: "admin@example.com"},
		},
	}
}

// GetUser fetches a user by ID
func (us *UserService) GetUser(userId string) (User, error) {
	user, exists := us.users[userId]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return user, nil
}
