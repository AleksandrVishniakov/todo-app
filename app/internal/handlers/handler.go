package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-app/app/internal/services"
	"todo-app/app/internal/services/auth"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/signin", h.signIn)
		auth.POST("/signup", h.signUp)
		auth.POST("/refresh", h.refreshTokens)
	}

	api := router.Group("/api", h.userAuthorization)
	{
		api.GET("/")
	}

	todo := router.Group("/todo", h.userAuthorization)
	{
		todo.GET("/", h.sayHello)
	}

	return router
}

func (h *Handler) refreshTokens(c *gin.Context) {
	var request = struct {
		RefreshToken string `json:"refreshToken"`
		Id           int    `json:"id"`
	}{}

	err := c.BindJSON(&request)
	if err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	responseTokens, err := h.service.Authorization.RefreshAccessToken(request.Id, request.RefreshToken)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, *responseTokens)
}

func (h *Handler) signIn(c *gin.Context) {
	var user auth.RequestUser
	err := c.BindJSON(&user)
	if err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	u, err := h.service.Authorization.SignInUser(&user)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	responseTokens, err := h.service.Authorization.SetTokensToUser(u.Id)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, *responseTokens)
}

func (h *Handler) signUp(c *gin.Context) {
	var newUser auth.RequestUser
	err := c.BindJSON(&newUser)
	if err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.CreateUser(&newUser)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	responseTokens, err := h.service.Authorization.SetTokensToUser(id)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, *responseTokens)
}

func (h *Handler) sayHello(c *gin.Context) {
	fmt.Fprint(c.Writer, "Heelo")
}
