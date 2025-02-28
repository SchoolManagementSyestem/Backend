package main

import (
	"errors"
	"sync"
)

type UserRepository struct {
	mu    sync.Mutex
	users map[string]*User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]*User),
	}
}

func (r *UserRepository) Save(user *User) (*User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.ID]; exists {
		return nil, errors.New("user already exists")
	}

	r.users[user.ID] = user
	return user, nil
}
