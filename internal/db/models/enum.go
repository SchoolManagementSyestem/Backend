package models

import (
	"database/sql/driver"
	"fmt"
)

// Role type (Enum)
type UserRole string

const (
	EnumAdmin   UserRole = "admin"
	EnumTeacher UserRole = "teacher"
	EnumStudent UserRole = "student"
	EnumStaff   UserRole = "staff"
	EnumParent  UserRole = "parent"
)

// Scan converts DB value to Role
func (r *UserRole) Scan(value interface{}) error {
	val, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid role value: %v", value)
	}
	*r = UserRole(val)
	return nil
}

// Value converts Role to DB-friendly format
func (r UserRole) Value() (driver.Value, error) {
	return string(r), nil
}

// Gender type (Enum)
type UserGender string

const (
	EnumMale   UserGender = "male"
	EnumFemale UserGender = "female"
	EnumOther  UserGender = "other"
)

// Scan converts DB value to Role
func (r *UserGender) Scan(value interface{}) error {
	val, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid role value: %v", value)
	}
	*r = UserGender(val)
	return nil
}

// Value converts Role to DB-friendly format
func (r UserGender) Value() (driver.Value, error) {
	return string(r), nil
}

// Status type (Enum)
type EnumStatus string

const (
	EnumActive    EnumStatus = "active"
	EnumInActive  EnumStatus = "inactive"
	EnumSuspended EnumStatus = "suspended"
)

// Scan converts DB value to Role
func (r *EnumStatus) Scan(value interface{}) error {
	val, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid role value: %v", value)
	}
	*r = EnumStatus(val)
	return nil
}

// Value converts Role to DB-friendly format
func (r EnumStatus) Value() (driver.Value, error) {
	return string(r), nil
}
