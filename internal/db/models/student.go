package models

import (
	"time"
)

type Student struct {
	Model
	UserID        string           `gorm:"not null;unique"`
	RollNumber    string           `gorm:"size:20;not null;unique"`
	ClassID       string           `gorm:"not null"`
	TeacherID     *string          `gorm:"default:null"`
	GuardianID    string           `gorm:"not null"`
	AdmissionDate time.Time        `gorm:"not null"`
	PaySchedule   uint             `gorm:"not null; default: 1"`
	Weaver        float64          `gorm:"type:numeric(8,2);default:0.00"`
	WeaverType    EnumDiscountType `gorm:"type:discount_type_enum;default:null"`
	User          *User            `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
