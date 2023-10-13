package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	timeZone := os.Getenv("DB_TIMEZONE")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + name + " port=" + port + " sslmode=disable TimeZone=" + timeZone
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		log.Fatal("Database Failed Connected")
	}

	DB = db
}