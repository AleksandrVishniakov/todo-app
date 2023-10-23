package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) userAuthorization(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		newResponseError(c, http.StatusUnauthorized, "empty auth header")
		c.Redirect(http.StatusUnauthorized, "http://localhost:8080/auth/signin")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newResponseError(c, http.StatusUnauthorized, "invalid auth header")
		c.Redirect(http.StatusUnauthorized, "http://localhost:8080/auth/signin")
		return
	}

	if len(headerParts[1]) == 0 {
		newResponseError(c, http.StatusUnauthorized, "token is empty")
		c.Redirect(http.StatusUnauthorized, "http://localhost:8080/auth/signin")
		return
	}

	tokenManager := h.service.Authorization.GetTokenManager()
	id, err := tokenManager.ParseAccessTokenWithId(headerParts[1])
	if err != nil {
		newResponseError(c, http.StatusUnauthorized, err.Error())
		c.Redirect(http.StatusUnauthorized, "http://localhost:8080/auth/signin")
		return
	}

	c.Set("id", id)
}
