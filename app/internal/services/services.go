package services

import (
	"time"
	"todo-app/app/internal/repositories"
	"todo-app/app/internal/services/auth"
	"todo-app/app/internal/services/auth/tokens"
)

type Service struct {
	Authorization
	ToDoManager
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{
		Authorization: auth.NewAuth(repo),
		ToDoManager:   newToDoManager(repo),
	}
}

type Authorization interface {
	CreateUser(user *auth.RequestUser) (int, error)
	GetTokenManager() auth.TokenManager
	SetTokensToUser(id int) (*tokens.ResponseTokens, error)
	RefreshAccessToken(id int, refreshToken string) (*tokens.ResponseTokens, error)
	SignInUser(user *auth.RequestUser) (*auth.ResponseUser, error)
}

type ToDoManager interface {
	GetAllToDosById(userId int) ([]ResponseToDo, error)
	AddToDo(userId int, todo RequestToDo) (ResponseToDo, error)
	UpdateToDoById(id int, messageText, messageColor string, messageDate time.Time) error
	DeleteToDoById(id int) error
}
