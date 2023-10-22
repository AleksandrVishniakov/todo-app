package handlers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/hello", h.sayHello)
	router.GET("/ping", h.ping)

	return router
}

func (h *Handler) sayHello(c *gin.Context) {
	fmt.Fprint(c.Writer, "Hello, World!")
}

func (h *Handler) ping(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Db-ping: %v", h.db.Ping())
}
