package main

import (
	"schoolManagementSystem/internal/db/models"
	pb "schoolManagementSystem/protos/auth"
)

func MapUserRoleToProto(role models.UserRole) pb.UserRole {
	switch role {
	case models.EnumStudent:
		return pb.UserRole_STUDENT
	case models.EnumTeacher:
		return pb.UserRole_TEACHER
	case models.EnumAdmin:
		return pb.UserRole_ADMIN
	case models.EnumStaff:
		return pb.UserRole_STAFF

	default:
		return pb.UserRole_ROLE_UNSPECIFIED
	}
}
func MapUserGenderToProto(role models.UserGender) pb.Gender {
	switch role {
	case models.EnumMale:
		return pb.Gender_MALE
	case models.EnumFemale:
		return pb.Gender_FEMALE
	case models.EnumOther:
		return pb.Gender_OTHER
	default:
		return pb.Gender(pb.Gender_GENDER_UNSPECIFIED)
	}
}
func MapStatusToProto(status models.EnumStatus) pb.Status {
	switch status {
	case models.EnumActive:
		return pb.Status_ACTIVE
	case models.EnumInActive:
		return pb.Status_INACTIVE
	case models.EnumSuspended:
		return pb.Status_SUSPENDED

	default:
		return pb.Status_STATUS_UNSPECIFIED
	}
}
