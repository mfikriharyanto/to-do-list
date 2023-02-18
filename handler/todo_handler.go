package handler

import (
	"net/http"
	"to-do-list/models"
	"to-do-list/services"

	"github.com/gin-gonic/gin"
)

func GetHelloWorld(context *gin.Context) {
	todos, _ := services.GetAllTodo()
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Hello World", "data": todos})
}

func GetTodos(context *gin.Context) {
	todos, _ := services.GetAllTodo()
	context.IndentedJSON(http.StatusOK, todos)
}

func GetTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := services.GetTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func PostTodo(context *gin.Context) {
	var createTodo models.CreateToDo

	err := context.ShouldBindJSON(&createTodo)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	todo, err := services.CreateTodo(createTodo)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusCreated, todo)
}