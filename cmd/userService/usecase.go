package main

import (
	"errors"
)

type UserUsecase struct {
	repo *UserRepository
}

func NewUserUsecase(repo *UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) CreateUser(name, email string) (*User, error) {
	if name == "" || email == "" {
		return nil, errors.New("name and email are required")
	}

	user := &User{
		ID:    "123", // Generate unique ID
		Name:  name,
		Email: email,
	}
	return u.repo.Save(user)
}
