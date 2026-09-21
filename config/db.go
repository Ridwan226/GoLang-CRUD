package config

import (
	"belajar/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(){
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("Please provide DATABASE_URL")
	}
	
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	
	err = database.AutoMigrate(&models.User{}, &models.Event{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	DB = database
	log.Println("Database connected")
}