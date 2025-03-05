package db

import (
	"log"
	"schoolManagementSystem/internal/db/models"

	"gorm.io/gorm"
)

func AutoMigrates(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Class{},
		&models.Parent{},
		&models.Qualification{},
		&models.Salary{},
		&models.Staff{},
		&models.Student{},
		&models.Teacher{},
		&models.User{})
	if err != nil {
		log.Fatalf("❌ Migration error: %v", err)
	}
}
