package main

import (
	"log"
	"to-do-list/handler"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=todolist port=5432"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection refused")
	}

	log.Print("Database connection success")

	router := gin.Default()
	router.GET("", handler.GetHelloWorld)

	router.Run(":8080")
}