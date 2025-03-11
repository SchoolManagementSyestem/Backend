// handler.go
package main

import (
	"context"
	authpb "schoolManagementSystem/protos/auth"
)

type AuthHandler struct {
	authpb.UnimplementedAuthServiceServer
	usecase *AuthUsecase // Change to a pointer
}

func NewAuthHandler(u *AuthUsecase) *AuthHandler { // Accept a pointer to AuthUsecase
	return &AuthHandler{usecase: u}
}

func (h *AuthHandler) LoginStaff(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	user, err := h.usecase.LoginStaff(req)
	if err != nil {
		return &authpb.LoginResponse{
			Status:  false,
			Message: err.Error(),
		}, nil
	}

	return user, nil
}
