package models

type Setting struct {
	Model
	Key   string `gorm:"type:varchar(50);not null"`
	Value string `gorm:"type:varchar(255);not null"`
}
