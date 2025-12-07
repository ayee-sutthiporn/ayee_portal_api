package database

import (
	"fmt"
	"log"

	"ayee-portal-backend/config"
	"ayee-portal-backend/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.AppConfig.DBHost,
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBName,
		config.AppConfig.DBPort,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Connected to Database successfully")

	// AutoMigrate
	log.Println("Running AutoMigrate...")
	err = DB.AutoMigrate(
		&models.User{},
		&models.Website{},
		&models.Category{},
		&models.AuditLog{},
		&models.SystemSettings{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
	log.Println("Database Migration completed")
}
