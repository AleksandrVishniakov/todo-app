package services

import (
	"time"
	"todo-app/app/internal/repositories"
)

var _ ToDoManager = &todoManager{}

type todoManager struct {
	repo *repositories.Repository
}

func newToDoManager(repo *repositories.Repository) *todoManager {
	return &todoManager{repo: repo}
}

func (t *todoManager) GetAllToDosById(userId int) ([]ResponseToDo, error) {
	ts, err := t.repo.ToDoManager.GetAllMessagesById(userId)
	if err != nil {
		return []ResponseToDo{}, err
	}

	var todos []ResponseToDo
	for _, todo := range ts {
		todos = append(todos, *mapToDoResponseFromEntity(&todo))
	}

	return todos, nil
}

func (t *todoManager) AddToDo(userId int, todo RequestToDo) (ResponseToDo, error) {
	td, err := t.repo.ToDoManager.AddMessage(userId, *mapToDoEntityFromRequest(&todo))
	return *mapToDoResponseFromEntity(&td), err
}

func (t *todoManager) UpdateToDoById(id int, messageText, messageColor string, messageDate time.Time) error {
	return t.repo.ToDoManager.UpdateMessageById(id, messageText, messageColor, messageDate)
}

func (t *todoManager) DeleteToDoById(id int) error {
	return t.repo.ToDoManager.DeleteMessageById(id)
}
