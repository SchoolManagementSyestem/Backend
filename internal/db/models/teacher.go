package models

type Teacher struct {
	Model
	Name     string `gorm:"type:varchar(255);not null"`
	Subject  string `gorm:"type:varchar(255);not null"`
	Salary   float64
	IsActive bool `gorm:"default:true"`
}
