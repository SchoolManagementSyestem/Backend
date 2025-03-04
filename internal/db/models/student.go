package models

import (
	"time"
)

type Student struct {
	Model
	UserID        uint      `gorm:"not null;unique"` // Foreign key for User
	RollNumber    string    `gorm:"size:20;not null;unique"`
	ClassID       uint      `gorm:"not null"`     // Assuming ClassID is a uint
	TeacherID     *uint     `gorm:"default:null"` // Assuming TeacherID is a uint pointer
	GuardianID    uint      `gorm:"not null"`     // Assuming GuardianID is a uint
	AdmissionDate time.Time `gorm:"not null"`
	PaySchedule   uint      `gorm:"not null; default: 1"` // Monthly, Quarterly, Annually
	Weaver        float64   `gorm:"type:numeric(8,2);default:0.00"`
	WeaverType    *string   `gorm:"size:20;default:null"`
	User          *User     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Use a pointer to User
}
