package models

import "time"

type Teacher struct {
	Model
	StaffID       string          `gorm:"not null;unique"`
	Department    string          `gorm:"type:varchar(30);default:null"`
	Subject       string          `gorm:"type:varchar(30);default:null"`
	Experience    string          `gorm:"type:varchar(100);default:null"`
	JoinedDate    time.Time       `gorm:"type:varchar(30);default:null"`
	Salary        float64         `gorm:"type:numeric(8,2);default:0.00"`
	Status        EnumStatus      `gorm:"default:active"`
	Staff         *Staff          `gorm:"foreignKey:StaffID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Qualification []Qualification `gorm:"foreignKey:TeacherID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
