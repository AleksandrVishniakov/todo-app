package repositories

import (
	"database/sql"
	"time"
)

type Repository struct {
	Authorization
	ToDoManager
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: newAuth(db),
		ToDoManager:   newToDo(db),
	}
}

type Authorization interface {
	GetUserById(id int) (UserEntity, error)
	GetUserByLogin(login string) (UserEntity, error)
	AddUser(u UserEntity) (UserEntity, error)
	GetRefreshTokenById(userId int) (TokenEntity, error)
	UpdateRefreshTokenById(t TokenEntity) error
}

type ToDoManager interface {
	GetAllMessagesById(userId int) ([]ToDoEntity, error)
	UpdateMessageById(id int, messageText, messageColor string, messageDate time.Time) error
	DeleteMessageById(id int) error
}
