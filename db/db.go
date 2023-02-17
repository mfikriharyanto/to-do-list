package db

import (
	"to-do-list/models"
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDatabase() (string, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	var err error
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", host, user, password, dbname, port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&models.ToDo{})

	if err != nil {
		return "", errors.New("database connection refused")
	}
	
	return "Database connection success", nil
}

func GetDB() *gorm.DB {
	return db
}