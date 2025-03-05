package models

type Class struct {
	Model
	Name      string `gorm:"size:100;not null;unique"`
	Section   string `gorm:"size:30;not null;"`
	teacherId string `gorm:"not null;unique"`
}
