package main

import (
	"log"
	"to-do-list/db"
	"to-do-list/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	message, err := db.ConnectDatabase()

	if err != nil {
		log.Printf("Error: %v", err)
	}

	log.Print(message)

	router := gin.Default()
	router.GET("", handler.GetHelloWorld)
	router.GET("/todos", handler.GetTodos)
	router.GET("/todos/:id", handler.GetTodo)
	router.POST("/todos", handler.PostTodo)

	router.Run(":8080")
}