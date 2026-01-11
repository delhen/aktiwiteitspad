package handlers

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"go_aktiwiteitspad/api/models"
)

var DB *gorm.DB

func connectDatabase(databaseUrl string) {
	var err any
	DB, err = gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}

func migrateDatabase() {
	if err := DB.AutoMigrate(&models.User{}, &models.ActivityTracker{}); err != nil {
		log.Fatal("Failed to migrate schema:", err)
	}
}

func InitializeDatabase() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	databaseUrl := os.Getenv("DB_URL")

	connectDatabase(databaseUrl)
	migrateDatabase()
}