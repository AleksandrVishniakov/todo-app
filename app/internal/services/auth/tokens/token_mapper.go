package tokens

import "todo-app/app/internal/repositories"

func MapRefreshTokenFromEntity(t *repositories.TokenEntity) *RefreshToken {
	return &RefreshToken{
		Val: t.RefreshToken,
		Exp: t.RefreshTokenExpiresAt,
	}
}

func MapTokenEntityFromRequest(t *RefreshToken) *repositories.TokenEntity {
	return &repositories.TokenEntity{
		RefreshToken:          t.Val,
		RefreshTokenExpiresAt: t.Exp,
	}
}
