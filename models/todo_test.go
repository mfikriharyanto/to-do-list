package models

import (
    "testing"
)

var (
	id uint = 1
	title string = "PPL Submission"
	completed bool = false

	todo = ToDo{
		ID: id,
		Title: title,
		Completed: completed,
	}

	createTodo = CreateToDo{
		Title: title,
		Completed: completed,
	}
)

func TestToDoID(t *testing.T) {
    if todo.ID != id {
		t.Errorf("got %d, wanted %d", todo.ID, id)
    }
}

func TestToDoTitle(t *testing.T) {
    if todo.Title != title {
		t.Errorf("got %s, wanted %s", todo.Title, title)
    }
}

func TestToDoCompleted(t *testing.T) {
    if todo.Completed != completed {
		t.Errorf("got %t, wanted %t", todo.Completed, completed)
    }
}

func TestCreateToDoTitle(t *testing.T) {
    if createTodo.Title != title {
		t.Errorf("got %s, wanted %s", createTodo.Title, title)
    }
}

func TestCreateToDoCompleted(t *testing.T) {
    if createTodo.Completed != completed {
		t.Errorf("got %t, wanted %t", createTodo.Completed, completed)
    }
}