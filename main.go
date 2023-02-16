package main

import (
	"log"
	"to-do-list/config"
	"to-do-list/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	message, err := config.ConnectDatabase()

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Print(message)

	router := gin.Default()
	router.GET("", handler.GetHelloWorld)

	router.Run(":8080")
}