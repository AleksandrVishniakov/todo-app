package services

import (
	"todo-app/app/internal/repositories"
)

func mapToDoEntityFromRequest(t *RequestToDo) *repositories.ToDoEntity {
	return &repositories.ToDoEntity{
		ToDoText:  t.ToDoText,
		ToDoDate:  t.ToDoDate,
		ToDoColor: t.ToDoColor,
	}
}

func mapToDoResponseFromEntity(t *repositories.ToDoEntity) *ResponseToDo {
	return &ResponseToDo{
		Id:        t.Id,
		ToDoText:  t.ToDoText,
		ToDoDate:  t.ToDoDate,
		ToDoColor: t.ToDoColor,
	}
}
