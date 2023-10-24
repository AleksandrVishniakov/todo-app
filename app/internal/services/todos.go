package services

import "time"

type RequestToDo struct {
	ToDoText  string    `json:"toDoText"`
	ToDoDate  time.Time `json:"toDoDate"`
	ToDoColor string    `json:"toDoColor"`
}

type ResponseToDo struct {
	Id        int       `json:"id"`
	ToDoText  string    `json:"toDoText"`
	ToDoDate  time.Time `json:"toDoDate"`
	ToDoColor string    `json:"toDoColor"`
}
