// handler.go
package main

import (
	"context"
	"fmt"
	pb "schoolManagementSystem/protos/user"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	usecase *UserUsecase // Change to a pointer
}

func NewUserHandler(u *UserUsecase) *UserHandler { // Accept a pointer to UserUsecase
	return &UserHandler{usecase: u}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := h.usecase.CreateUser(req.Name, req.Email)
	if err != nil {
		return &pb.CreateUserResponse{
			Status:  false,
			Message: "User already exists",
		}, nil
	}

	fmt.Println("User created successfully")

	return &pb.CreateUserResponse{
		Status:  true,
		Message: "User created successfully",
		Data: &pb.CreateUserResponseData{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}
