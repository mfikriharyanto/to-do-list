package main

import (
	"to-do-list/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("", handler.GetHelloWorld)

	router.Run(":8080")
}