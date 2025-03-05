package models

type Parent struct {
	Model
	UserID     uint           `gorm:"not null;unique"`
	Reletion   EnumParentType `gorm:"type:parent_type_enum;not null;"`
	Income     float64        `gorm:"type:numeric(8,2);default:0.00"`
	IncomeType EnumIncomeType `gorm:"type:income_type_enum;default:null"`
	Status     EnumStatus     `gorm:"default:active"`
	User       *User          `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
