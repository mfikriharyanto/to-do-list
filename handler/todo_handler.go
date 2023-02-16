package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHelloWorld(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Hello World"})
}
