package models

type Salary struct {
	Model
	StaffID   string                `gorm:"not null;unique"`
	Amount    float64               `gorm:"type:numeric(8,2);default:0.00"`
	NetSalary float64               `gorm:"type:numeric(8,2);default:0.00"`
	Month     string                `gorm:"not null"`
	Year      string                `gorm:"not null"`
	PayDate   string                `gorm:"not null"`
	Status    EnumTransactionStatus `gorm:"default:pending"`
	Staff     *Staff                `gorm:"foreignKey:StaffID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
