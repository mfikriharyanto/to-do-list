package models

type ToDo struct {
	ID     		uint	`json:"id"`
	Title  		string	`json:"title"`
	Completed 	bool	`json:"completed"`
}

type CreateToDo struct {
	Title		string	`json:"title" binding:"required"`
	Completed 	bool 	`json:"completed"`
}