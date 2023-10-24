package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) userAuthorization(c *gin.Context) {
	var id int
	tokenManager := h.service.Authorization.GetTokenManager()

	header := c.GetHeader("Authorization")
	if header == "" {
		newResponseError(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newResponseError(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		newResponseError(c, http.StatusUnauthorized, "token is empty")
		return
	}

	id, err := tokenManager.ParseAccessTokenWithId(headerParts[1])
	if err != nil {
		newResponseError(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("id", id)
}

func (h *Handler) authCookieEncoder() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth, err := h.store.Get(c.Request, "auth-cookie")
		if err != nil {
			newResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		var id int
		var refreshToken string
		var idVal = auth.Values["id"]
		if idVal == nil {
			id = 0
		} else {
			id = idVal.(int)
		}

		var refreshTokenVal = auth.Values["refresh_token"]
		if refreshTokenVal == nil {
			refreshToken = ""
		} else {
			refreshToken = refreshTokenVal.(string)
		}

		c.Set("id", id)
		c.Set("refresh_token", refreshToken)
	}
}
