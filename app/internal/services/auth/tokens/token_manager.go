package tokens

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"log"
	"math/rand"
	"time"
)

var (
	tokenCfg = tokenConfigs{}
)

func init() {
	//viper initialization
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Viper init error: %s", err.Error())
	}

	tokenCfg.refreshTokenTTL = viper.GetInt64("tokens.refresh.ttl")
	tokenCfg.accessTokenTTL = viper.GetInt64("tokens.access.ttl")
}

type JWTokenManager struct {
	signature string
}

type tokenConfigs struct {
	accessTokenTTL  int64 `yaml:"accessTokenTTL"`
	refreshTokenTTL int64 `yaml:"refreshTokenTTL"`
}

func NewJWTokenManager(signature string) *JWTokenManager {
	return &JWTokenManager{
		signature: signature,
	}
}

func (m *JWTokenManager) GenerateRefreshToken() *RefreshToken {

	b := make([]byte, 64)

	s := rand.NewSource(time.Now().Unix() * rand.Int63())
	r := rand.New(s)

	_, err := r.Read(b)
	if err != nil {
		log.Fatalf("Refresh token generation: %s", err.Error())
	}

	return NewRefreshToken(fmt.Sprintf("%x", b), time.Now().Unix()+tokenCfg.refreshTokenTTL)
}

func (m *JWTokenManager) GenerateAccessToken(id int) string {
	token := NewAccessToken(
		accessTokenBody{
			id:  id,
			exp: time.Now().Unix() + tokenCfg.accessTokenTTL,
		},
	)

	fmt.Printf("Now:%d\tToken:%d\n", time.Now().Unix(), token.body.exp)

	tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  token.body.id,
		"exp": token.body.exp,
	})

	tokenString, err := tokenJWT.SignedString([]byte(m.signature))
	if err != nil {
		log.Fatalf("Access token generation: %s", err.Error())
	}
	return tokenString
}

func (m *JWTokenManager) ParseAccessTokenWithId(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token method")
		}

		return []byte(m.signature), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp := int64(claims["exp"].(float64))
		if err != nil {
			return 0, err
		}

		if IsExpired(exp) {
			return 0, errors.New("token is expired")
		}

		return int(claims["id"].(float64)), nil
	}

	return 0, errors.New("unexpected error")
}
