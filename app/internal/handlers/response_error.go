package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type responseError struct {
	Message   string    `json:"message"`
	Code      int       `json:"code"`
	Timestamp time.Time `json:"timestamp"`
}

func newResponseError(c *gin.Context, code int, message string) {
	log.Printf("%d\t%s\n", code, message)
	c.AbortWithStatusJSON(code, responseError{
		Message:   message,
		Code:      code,
		Timestamp: time.Now(),
	})
}
