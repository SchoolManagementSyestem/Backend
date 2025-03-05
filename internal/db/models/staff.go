package models

type Staff struct {
	Model
	UserID   uint              `gorm:"not null;unique"`
	Position EnumStaffPosition `gorm:"type:staff_position_enum;not null;"`
	User     *User             `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
