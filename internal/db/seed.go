package db

import (
	"log"
	"schoolManagementSystem/internal/db/models"
	"schoolManagementSystem/pkg/helpers"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) {
	tenantID, err := uuid.Parse("61b37884-8c0d-4d1c-ba89-68bcb0f4bd60")
	if err != nil {
		log.Fatal("Invalid UUID for TenantId:", err)
	}
	seedUsers(db, tenantID)
}
func SeedSettings(db *gorm.DB) {
	tenantID, err := uuid.Parse("61b37884-8c0d-4d1c-ba89-68bcb0f4bd60")
	if err != nil {
		log.Fatal("Invalid UUID for TenantId:", err)
	}
	seedSettings(db, tenantID)
}

// SeedSettings inserts default settings into the database
func seedSettings(db *gorm.DB, tenantID uuid.UUID) {
	settings := []models.Setting{
		{Model: models.Model{TenantId: tenantID}, Key: "roll_format", Value: "CUS=ABC, YEAR, SHORT_YEAR, MONTH, DAY, SEQ=0001, RANDOM=0000, CLASS_CODE, SECTION_CODE"},
		{Model: models.Model{TenantId: tenantID}, Key: "app_name", Value: "School Management System"},
		{Model: models.Model{TenantId: tenantID}, Key: "site_url", Value: "https://school.example.com"},
		{Model: models.Model{TenantId: tenantID}, Key: "timezone", Value: "Asia/Dhaka"},
		{Model: models.Model{TenantId: tenantID}, Key: "currency", Value: "BDT"},
		{Model: models.Model{TenantId: tenantID}, Key: "currency_symbol", Value: "৳"},
		{Model: models.Model{TenantId: tenantID}, Key: "date_format", Value: "d-m-Y"},
		{Model: models.Model{TenantId: tenantID}, Key: "time_format", Value: "h:i A"},
		{Model: models.Model{TenantId: tenantID}, Key: "datetime_format", Value: "d-m-Y h:i A"},
		{Model: models.Model{TenantId: tenantID}, Key: "language", Value: "en"},
		{Model: models.Model{TenantId: tenantID}, Key: "locale", Value: "en_US"},
		{Model: models.Model{TenantId: tenantID}, Key: "last_uid", Value: "1000"},
		{Model: models.Model{TenantId: tenantID}, Key: "roll_format_seq", Value: "1000"},
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
}

func seedUsers(db *gorm.DB, tenantID uuid.UUID) {
	password, err := helpers.HashPassword("password")
	if err != nil {
		log.Fatal("Failed to hash password")
	}
	users := []models.User{
		{Model: models.Model{TenantId: tenantID}, Password: password, UID: 1, FirstName: "Admin", Role: models.EnumAdmin, DateOfBirth: "06-06-1990", Gender: models.EnumMale, Status: models.EnumActive},
	}

	for _, user := range users {
		var existing models.User

		if err := db.Where("uid = ?", user.UID).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&user).Error; err != nil {
					log.Printf("Failed to seed user (%d): %v", user.UID, err)
				}
			} else {
				log.Printf("Error checking user (%d): %v", user.UID, err)
			}
		}
	}
}
