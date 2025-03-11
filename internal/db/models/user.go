package models

type User struct {
	Model
	UID            uint       `gorm:"uniqueIndex; autoIncrement; <-:create; not null;"`
	FirstName      string     `gorm:"type:varchar(50);not null"`
	LastName       string     `gorm:"type:varchar(50);default:null"`
	Email          string     `gorm:"type:varchar(50);default:null"`
	Phone          string     `gorm:"type:varchar(50);default:null"`
	Password       string     `gorm:"type:varchar(255);not null"`
	Role           UserRole   `gorm:"type:role_enum;not null;default:'student'"`
	DateOfBirth    string     `gorm:"type:varchar(50);default:null"`
	Gender         UserGender `gorm:"type:gender_enum;not null;default:null"`
	Address        string     `gorm:"type:varchar(255);default:null"`
	ProfilePicture string     `gorm:"type:varchar(255);default:null"`
	Status         EnumStatus `gorm:"type:status_enum;not null;default:'active'"`
	Student        *Student   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Staff          *Staff     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Parent         *Parent    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
