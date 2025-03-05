package models

type Class struct {
	Model
	Name      string     `gorm:"size:100;not null;unique"`
	Section   string     `gorm:"size:30;not null;"`
	Capacity  uint       `gorm:"not null;default:0"`
	Status    EnumStatus `gorm:"default:active"`
	TeacherId string     `gorm:"not null;unique"`
	Teacher   *Teacher   `gorm:"foreignKey:TeacherId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
