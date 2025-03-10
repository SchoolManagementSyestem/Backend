package main

import (
	"errors"
	"schoolManagementSystem/pkg/helpers"
	pb "schoolManagementSystem/protos/auth"
)

type AuthUsecase struct {
	repo *AuthRepository
}

func NewAuthUsecase(repo *AuthRepository) *AuthUsecase {
	return &AuthUsecase{repo: repo}
}

func (u *AuthUsecase) LoginStaff(uid, password string) (*pb.LoginResponse, error) {
	// Fetch user by ID
	user, err := u.repo.GetUserByID(uid)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Validate password
	if err := helpers.CheckPasswordHash(password, user.Password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT tokens
	accessToken, err := helpers.GenerateAccessToken(user.ID, string(user.Role), user.TenantId)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	refreshToken, err := helpers.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	// Create token response
	token := &pb.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	// Create login response data
	loginResponseData := &pb.LoginResponseData{
		Id:             user.ID.String(),
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		Phone:          user.Phone,
		Role:           MapUserRoleToProto(user.Role),
		DateOfBirth:    user.DateOfBirth,
		Gender:         MapUserGenderToProto(user.Gender),
		Address:        user.Address,
		ProfilePicture: user.ProfilePicture,
		Status:         MapStatusToProto(user.Status),
		CreatedAt:      user.CreatedAt.String(),
		UpdatedAt:      user.UpdatedAt.String(),
		DeletedAt:      user.DeletedAt.Time.String(),
		Token:          token,
	}

	// Returning the response
	return &pb.LoginResponse{
		Status:  true,
		Message: "Login successful",
		Data:    loginResponseData,
	}, nil
}
