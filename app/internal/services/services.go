package services

import (
	"todo-app/app/internal/repositories"
	"todo-app/app/internal/services/auth"
	"todo-app/app/internal/services/auth/tokens"
)

type Service struct {
	Authorization
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{
		Authorization: auth.NewAuth(repo),
	}
}

type Authorization interface {
	CreateUser(user *auth.RequestUser) (int, error)
	GetTokenManager() auth.TokenManager
	SetTokensToUser(id int) (*tokens.ResponseTokens, error)
	RefreshAccessToken(id int, refreshToken string) (*tokens.ResponseTokens, error)
	SignInUser(user *auth.RequestUser) (*auth.ResponseUser, error)
}
