package db

import (
	"fmt"
	"log"
	"os"
	"schoolManagementSystem/internal/db/models"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB   *gorm.DB
	once sync.Once
)

// InitDB initializes the PostgreSQL database connection
func InitPGDB() *gorm.DB {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Dhaka",
			getEnv("DB_HOST", "localhost"),
			getEnv("DB_USER", "postgres"),
			getEnv("DB_PASS", "password"),
			getEnv("DB_NAME", "school_db"),
			getEnv("DB_PORT", "5432"),
		)

		var err error
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatalf("❌ Failed to connect to database: %v", err)
		}

		// Apply Migrations
		autoMigrate(DB)
	})

	return DB
}

// AutoMigrate applies database migrations
func autoMigrate(db *gorm.DB) {
	createEnum(db, "role_enum", "'admin', 'teacher', 'student', 'staff', 'parent'")
	createEnum(db, "parent_type_enum", "'father', 'mother', 'student', 'guardian', 'relative', 'other'")
	createEnum(db, "income_type_enum", "'business', 'job', 'agriculture', 'pension', 'foreigner', 'other'")
	createEnum(db, "grade_point_enum", "'gpa', 'cgpa', 'marks', 'percent'")
	createEnum(db, "status_enum", "'active', 'inactive', 'suspended'")
	createEnum(db, "gender_enum", "'male', 'female', 'other'")
	createEnum(db, "staff_position_enum", "'principle', 'accountant', 'librarian', 'receptionist', 'clerk', 'peon', 'driver', 'security', 'cleaner', 'cook', 'shopkeeper', 'wathman', 'electrician', 'sweeper', 'nurse', 'care_taker', 'lab_assistant', 'other'")
	createEnum(db, "qualification_type_enum", "'ssc', 'hsc', 'diploma', 'bachelor', 'master', 'phd', 'other'")
	createEnum(db, "discount_type_enum", "'flat', 'percentage'")

	err := db.AutoMigrate(&models.Parent{}, &models.Qualification{}, &models.Staff{}, &models.Student{}, &models.Teacher{}, &models.User{})
	if err != nil {
		log.Fatalf("❌ Migration error: %v", err)
	}
	fmt.Println("✅ Migrations applied successfully!")
}

// Helper function to get environment variables with a default fallback
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// createEnum ensures an ENUM type exists before using it
func createEnum(db *gorm.DB, enumName, values string) {
	query := fmt.Sprintf("DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = '%s') THEN CREATE TYPE %s AS ENUM (%s); END IF; END $$;", enumName, enumName, values)
	db.Exec(query)
}
