package models

type Qualification struct {
	Model
	TeacherID string                `gorm:"not null;unique"`
	Type      EnumQualificationType `gorm:"type:qualification_type_enum;not null;"`
	Institute string                `gorm:"type:varchar(100);not null"`
	Year      string                `gorm:"type:varchar(4);not null"`
	Grade     float64               `gorm:"type:numeric(8,2);default:0.00"`
	GradeType EnumGradePoint        `gorm:"type:grade_point_enum;not null;default:null"`
	IsActive  bool                  `gorm:"default:true"`
	Teacher   *Teacher              `gorm:"foreignKey:TeacherID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
