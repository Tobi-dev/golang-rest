package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func connect() {
	log.Println("Initializing config file...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading .env file!")
	}

	driver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")

	log.Println("Config successfully initialized.")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, database, password)

	conn, err := gorm.Open(driver, dbURI)

	if err != nil {
		log.Fatal(err)
	}

	db = conn
}

func GetDB() *gorm.DB {
	connect()
	return db;
}