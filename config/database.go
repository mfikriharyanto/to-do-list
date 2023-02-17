package config

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (string, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", host, user, password, dbname, port)
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Print(dsn)

	if err != nil {
		return "", errors.New("database connection refused")
	}
	
	return "Database connection success", nil
}