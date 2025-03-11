package models

import "time"

type Staff struct {
	Model
	UserID        string            `gorm:"not null;unique"`
	Position      EnumStaffPosition `gorm:"type:staff_position_enum;not null;"`
	Salary        float64           `gorm:"type:numeric(8,2);default:0.00"`
	Status        EnumStatus        `gorm:"default:active"`
	User          *User             `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Experience    string            `gorm:"type:varchar(100);default:null"`
	Qualification []Qualification   `gorm:"foreignKey:TeacherID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	JoinedDate    time.Time         `gorm:"type:varchar(30);default:null"`
	Teacher       *Teacher          `gorm:"foreignKey:StaffID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
