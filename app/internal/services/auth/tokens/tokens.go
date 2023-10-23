package tokens

import (
	"time"
)

type ResponseTokens struct {
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
}

func NewResponseTokens(refreshToken, accessToken string) *ResponseTokens {
	return &ResponseTokens{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}
}

type RefreshToken struct {
	Val string
	Exp int64
}

func NewRefreshToken(val string, exp int64) *RefreshToken {
	return &RefreshToken{
		Val: val,
		Exp: exp,
	}
}

type AccessToken struct {
	body accessTokenBody
}

type accessTokenBody struct {
	id  int
	exp int64
}

func NewAccessToken(body accessTokenBody) *AccessToken {
	return &AccessToken{
		body: body,
	}
}

func IsExpired(exp int64) bool {
	return time.Now().Unix() > exp
}
