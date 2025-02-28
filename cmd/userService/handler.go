// handler.go
package main

import (
	"context"
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
		return nil, err
	}

	return &pb.CreateUserResponse{
		Status:  "success",
		Message: "User created successfully",
		Data: &pb.CreateUserResponseData{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}
