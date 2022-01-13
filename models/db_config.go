package models

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

var DB *Database

func ConnectDatabase() {
	err := godotenv.Load("service.env")
	if err != nil {
		log.Fatalf("Error when loading env variables. Err: %s", err)
	}

	db_name := os.Getenv("DB_NAME")
	database, err := gorm.Open(sqlite.Open(db_name), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	database.AutoMigrate(&User{})
	DB.DB = database
}
