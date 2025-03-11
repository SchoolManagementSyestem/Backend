package models

type Teacher struct {
	Model
	StaffID    string `gorm:"not null;unique"`
	Department string `gorm:"type:varchar(30);default:null"`
	Section    string `gorm:"type:varchar(30);default:null"`
	Subject    string `gorm:"type:varchar(30);default:null"`
	Staff      *Staff `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
