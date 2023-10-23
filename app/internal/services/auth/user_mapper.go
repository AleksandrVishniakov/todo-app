package auth

import (
	"todo-app/app/internal/repositories"
	"todo-app/app/internal/services/password"
)

func mapResponseUserFromEntity(user *repositories.UserEntity) *ResponseUser {
	return &ResponseUser{
		Id:    user.Id,
		Login: user.Login,
	}
}

func mapEntityUserFromRequest(user *RequestUser) *repositories.UserEntity {
	salt := password.GenerateSalt()
	return &repositories.UserEntity{
		Login:    user.Login,
		Password: password.GeneratePasswordHash(user.Password, salt),
		Salt:     salt,
	}
}
