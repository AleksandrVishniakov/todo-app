package repositories

import "time"

type ToDoEntity struct {
	Id        int
	UserId    int
	ToDoText  string
	ToDoDate  time.Time
	ToDoColor string
}
