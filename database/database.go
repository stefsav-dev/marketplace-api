package database

import (
	"log"
	"os"
	"test_marketplace/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=marketplace port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	}
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	err = database.AutoMigrate(&models.User{}, &models.Product{}, &models.Transaction{})

	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}

	DB = database
	log.Println("Connected to database")
}
