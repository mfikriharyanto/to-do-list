package config

import (
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/spf13/viper"
)

func ConnectDatabase() (string, error) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		return "", errors.New("failed while reading config file")
	}

	host := viper.Get("DB_HOST").(string)
	user := viper.Get("DB_USER").(string)
	password := viper.Get("DB_PASSWORD").(string)
	dbname := viper.Get("DB_NAME").(string)
	port := viper.Get("DB_PORT").(string)

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", host, user, password, dbname, port)
	_, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return "", errors.New("database connection refused")
	}
	
	return "Database connection success", nil
}