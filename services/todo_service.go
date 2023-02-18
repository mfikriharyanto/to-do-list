package services

import (
	"errors"
	"to-do-list/db"
	"to-do-list/models"
)

func GetAllTodo() (*[]models.ToDo, error) {
	var todos []models.ToDo
	
	db := db.GetDB()

	if result := db.Find(&todos); result.Error != nil {
		return nil, errors.New("todos not found")
	}
	
	return &todos, nil
}

func GetTodoById(id string) (*models.ToDo, error) {
	var todo models.ToDo
	
	db := db.GetDB()

	if result := db.First(&todo, id); result.Error != nil {
		return nil, errors.New("todo not found")
	}
	
	return &todo, nil
}

func CreateTodo(createTodo models.CreateToDo) (*models.ToDo, error) {
	todo := models.ToDo{
		Title: createTodo.Title,
		Completed: createTodo.Completed,
	}
	
	db := db.GetDB()

	if result := db.Create(&todo); result.Error != nil {
		return nil, errors.New("failed to create todo")
	}
	
	return &todo, nil
}