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

/***************** Status Enum Start *****************/

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

/***************** Status Enum End *****************/

/***************** Position Enum Start *****************/

type EnumStaffPosition string

const (
	EnumStaffPositionPrinciple    EnumStaffPosition = "principle"
	EnumStaffPositionAccountant   EnumStaffPosition = "accountant"
	EnumStaffPositionLibrarian    EnumStaffPosition = "librarian"
	EnumStaffPositionReceptionist EnumStaffPosition = "receptionist"
	EnumStaffPositionClerk        EnumStaffPosition = "clerk"
	EnumStaffPositionPeon         EnumStaffPosition = "peon"
	EnumStaffPositionDriver       EnumStaffPosition = "driver"
	EnumStaffPositionSecurity     EnumStaffPosition = "security"
	EnumStaffPositionCleaner      EnumStaffPosition = "cleaner"
	EnumStaffPositionCook         EnumStaffPosition = "cook"
	EnumStaffShopkeeper           EnumStaffPosition = "shopkeeper"
	EnumStaffWatchman             EnumStaffPosition = "watchman"
	EnumStaffElectrician          EnumStaffPosition = "electrician"
	EnumStaffPositionNurse        EnumStaffPosition = "nurse"
	EnumStaffPositionCareTaker    EnumStaffPosition = "care_taker"
	EnumStaffPositionLabAssistant EnumStaffPosition = "lab_assistant"
	EnumStaffPositionSweeper      EnumStaffPosition = "sweeper"
	EnumStaffPositionOther        EnumStaffPosition = "other"
)

// Scan converts DB value to Role
func (r *EnumStaffPosition) Scan(value interface{}) error {
	val, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid role value: %v", value)
	}
	*r = EnumStaffPosition(val)
	return nil
}

// Value converts Role to DB-friendly format
func (r EnumStaffPosition) Value() (driver.Value, error) {
	return string(r), nil
}

/***************** Position Enum End *****************/

/***************** GradePoint Enum Start *****************/

type EnumGradePoint string

const (
	EnumGradePointGPA     EnumGradePoint = "gpa"
	EnumGradePointCGPA    EnumGradePoint = "cgpa"
	EnumGradePointMarks   EnumGradePoint = "marks"
	EnumGradePointPercent EnumGradePoint = "percent"
)

// Scan converts DB value to Role
func (r *EnumGradePoint) Scan(value interface{}) error {
	val, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid role value: %v", value)
	}
	*r = EnumGradePoint(val)
	return nil
}

// Value converts Role to DB-friendly format
func (r EnumGradePoint) Value() (driver.Value, error) {
	return string(r), nil
}

/***************** GradePoint Enum End *****************/

/***************** Discount Enum Start *****************/

type EnumDiscountType string

const (
	EnumDiscountTypeFixed      EnumDiscountType = "fixed"
	EnumDiscountTypePercentage EnumDiscountType = "percentage"
)

// Scan converts DB value to Role
func (r *EnumDiscountType) Scan(value interface{}) error {
	val, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid role value: %v", value)
	}
	*r = EnumDiscountType(val)
	return nil
}

// Value converts Role to DB-friendly format
func (r EnumDiscountType) Value() (driver.Value, error) {
	return string(r), nil
}

/***************** Discount Enum End *****************/

/***************** EnumParentType Enum Start *****************/

type EnumParentType string

const (
	EnumParentTypeFather   EnumParentType = "father"
	EnumParentTypeMother   EnumParentType = "mother"
	EnumParentTypeGuardian EnumParentType = "guardian"
	EnumParentTypeRelative EnumParentType = "relative"
	EnumParentTypeOther    EnumParentType = "other"
)

// Scan converts DB value to Role
func (r *EnumParentType) Scan(value interface{}) error {
	val, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid role value: %v", value)
	}
	*r = EnumParentType(val)
	return nil
}

// Value converts Role to DB-friendly format
func (r EnumParentType) Value() (driver.Value, error) {
	return string(r), nil
}

/***************** EnumParentType Enum End *****************/

/***************** EnumIncomeType Enum Start *****************/

type EnumIncomeType string

const (
	EnumIncomeTypeBusiness    EnumIncomeType = "business"
	EnumIncomeTypeJob         EnumIncomeType = "job"
	EnumIncomeTypeAgriculture EnumIncomeType = "agriculture"
	EnumIncomeTypeTution      EnumIncomeType = "tuition"
	EnumIncomeTypePension     EnumIncomeType = "pension"
	EnumIncomeTypeForeigner   EnumIncomeType = "foreigner"
	EnumIncomeTypeOther       EnumIncomeType = "other"
)

// Scan converts DB value to Role
func (r *EnumIncomeType) Scan(value interface{}) error {
	val, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid role value: %v", value)
	}
	*r = EnumIncomeType(val)
	return nil
}

// Value converts Role to DB-friendly format
func (r EnumIncomeType) Value() (driver.Value, error) {
	return string(r), nil
}

/***************** EnumIncomeType Enum End *****************/

/***************** EnumQualificationType Enum Start *****************/

type EnumQualificationType string

const (
	EnumQualificationTypeJsc      EnumQualificationType = "jsc"
	EnumQualificationTypeSSC      EnumQualificationType = "ssc"
	EnumQualificationTypeHSC      EnumQualificationType = "hsc"
	EnumQualificationTypeDiploma  EnumQualificationType = "diploma"
	EnumQualificationTypeBachelor EnumQualificationType = "bachelor"
	EnumQualificationTypeMaster   EnumQualificationType = "master"
	EnumQualificationTypePhD      EnumQualificationType = "phd"
	EnumQualificationTypeOther    EnumQualificationType = "other"
)

// Scan converts DB value to Role
func (r *EnumQualificationType) Scan(value interface{}) error {
	val, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid role value: %v", value)
	}
	*r = EnumQualificationType(val)
	return nil
}

// Value converts Role to DB-friendly format
func (r EnumQualificationType) Value() (driver.Value, error) {
	return string(r), nil
}

/***************** EnumQualificationType Enum End *****************/
/***************** EnumTransactionStatus Enum Start *****************/

type EnumTransactionStatus string

const (
	EnumTransactionStatusPending EnumTransactionStatus = "pending"
	EnumTransactionStatusSuccess EnumTransactionStatus = "success"
	EnumTransactionStatusFailed  EnumTransactionStatus = "failed"
	EnumTransactionStatusRefund  EnumTransactionStatus = "refund"
	EnumTransactionStatusCancel  EnumTransactionStatus = "cancel"
	EnumTransactionStatusHold    EnumTransactionStatus = "hold"
	EnumTransactionStatusCharge  EnumTransactionStatus = "charge"
)

// Scan converts DB value to Role
func (r *EnumTransactionStatus) Scan(value interface{}) error {
	val, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid role value: %v", value)
	}
	*r = EnumTransactionStatus(val)
	return nil
}

// Value converts Role to DB-friendly format
func (r EnumTransactionStatus) Value() (driver.Value, error) {
	return string(r), nil
}

/***************** EnumTransactionStatus Enum End *****************/
