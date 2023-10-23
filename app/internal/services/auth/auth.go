package auth

import (
	"errors"
	"log"
	"os"
	"todo-app/app/internal/repositories"
	"todo-app/app/internal/services/auth/tokens"
	"todo-app/app/internal/services/password"
)

type Auth struct {
	TokenManager
	repo *repositories.Repository
}

func NewAuth(repo *repositories.Repository) *Auth {
	signature, ok := os.LookupEnv("TOKEN_SIGNATURE")
	if !ok {
		log.Fatal("no signature found")
	}

	return &Auth{
		TokenManager: tokens.NewJWTokenManager(signature),
		repo:         repo,
	}
}

func (a *Auth) CreateUser(user *RequestUser) (int, error) {
	u, err := a.repo.Authorization.AddUser(*mapEntityUserFromRequest(user))
	return u.Id, err
}

func (a *Auth) SignInUser(user *RequestUser) (*ResponseUser, error) {
	userLogged, err := a.repo.Authorization.GetUserByLogin(user.Login)
	if err != nil {
		return nil, err
	}
	if password.GeneratePasswordHash(user.Password, userLogged.Salt) != userLogged.Password {
		return nil, errors.New("incorrect password")
	}
	return mapResponseUserFromEntity(&userLogged), nil
}

func (a *Auth) GetTokenManager() TokenManager {
	return a.TokenManager
}

func (a *Auth) SetTokensToUser(id int) (*tokens.ResponseTokens, error) {
	refreshToken := a.TokenManager.GenerateRefreshToken()
	rTokens := tokens.NewResponseTokens(
		refreshToken.Val,
		a.TokenManager.GenerateAccessToken(id),
	)

	tEntity := *tokens.MapTokenEntityFromRequest(refreshToken)
	tEntity.UserId = id
	err := a.repo.Authorization.UpdateRefreshTokenById(
		tEntity,
	)

	return rTokens, err
}

func (a *Auth) RefreshAccessToken(id int, refreshToken string) (*tokens.ResponseTokens, error) {
	tEntity, err := a.repo.GetRefreshTokenById(id)
	token := tokens.MapRefreshTokenFromEntity(&tEntity)
	if err != nil {
		return nil, err
	}

	if token.Val != refreshToken {
		return nil, errors.New("incorrect refresh token")
	}

	if tokens.IsExpired(token.Exp) {
		return nil, errors.New("token is expired")
	}

	return a.SetTokensToUser(id)
}

type TokenManager interface {
	GenerateRefreshToken() *tokens.RefreshToken
	GenerateAccessToken(id int) string
	ParseAccessTokenWithId(tokenString string) (int, error)
}
