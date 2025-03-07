package db

import (
	"log"
	"schoolManagementSystem/internal/db/models"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	seedSettings(db)
}

// SeedSettings inserts default settings into the database
func seedSettings(db *gorm.DB) {
	settings := []models.Setting{
		{Key: "roll_format", Value: "CUS=ABC, YEAR, SHORT_YEAR, MONTH, DAY, SEQ=0001, RANDOM=0000, CLASS_CODE, SECTION_CODE"},
		{Key: "app_name", Value: "School Management System"},
		{Key: "site_name", Value: "School Management System"},
		{Key: "site_url", Value: "https://school.example.com"},
		{Key: "admin_email", Value: "admin@example.com"},
		{Key: "timezone", Value: "Asia/Dhaka"},
		{Key: "currency", Value: "BDT"},
	}

	// Insert settings, avoiding duplicates
	for _, setting := range settings {
		var existing models.Setting
		if err := db.Where("key = ?", setting.Key).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&setting).Error; err != nil {
					log.Printf("Failed to seed setting (%s): %v", setting.Key, err)
				}
			} else {
				log.Printf("Error checking setting (%s): %v", setting.Key, err)
			}
		}
	}

	log.Println("✅ Settings seeding completed!")
}
