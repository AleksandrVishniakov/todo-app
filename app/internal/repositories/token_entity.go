package repositories

type TokenEntity struct {
	UserId                int
	RefreshToken          string
	RefreshTokenExpiresAt int64
}
