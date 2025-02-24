// internal/graphql/user/service.go
package user

import "errors"

// User struct
type User struct {
	ID    int
	Name  string
	Email string
}

// Simulated in-memory user storage
var users = []User{
	{ID: 1, Name: "John Doe", Email: "john@example.com"},
}

// FindUserByID fetches a user by ID
func FindUserByID(id int) (*User, error) {
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

// CreateUser adds a new user
func CreateUser(name, email string) User {
	newUser := User{ID: len(users) + 1, Name: name, Email: email}
	users = append(users, newUser)
	return newUser
}
