package db

import (
	"fmt"
	"log"
	"os"
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
	// Create ENUM types
	CreateEnums(db)

	// Apply migrations
	AutoMigrates(db)

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
