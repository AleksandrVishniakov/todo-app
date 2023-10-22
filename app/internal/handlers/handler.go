package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/hello", h.sayHello)

	return router
}

func (h *Handler) sayHello(c *gin.Context) {
	fmt.Fprint(c.Writer, "Hello, World!")
}
