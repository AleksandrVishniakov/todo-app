package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
	"os"
	"todo-app/app/internal/services"
	"todo-app/app/internal/services/auth"
)

type Handler struct {
	service *services.Service
	store   *sessions.CookieStore
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{
		service: service,
		store:   sessions.NewCookieStore([]byte(os.Getenv("COOKIE_STORE_KEY"))),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(h.authCookieEncoder())

	//router.Static("/", "web/build/")
	router.LoadHTMLFiles("web/build/index.html", "web/build/sign-in.html", "web/build/sign-up.html")

	authorization := router.Group("/auth")
	{
		authorization.POST("/signin", h.signIn)
		authorization.POST("/signup", h.signUp)
		authorization.GET("/refresh", h.refreshTokens)
	}

	api := router.Group("/api", h.userAuthorization)
	{
		api.GET("/todos")
		api.DELETE("/todo")
		api.PUT("/todo")
		api.POST("/todo")
	}

	todo := router.Group("/todo", h.userAuthorization)
	{
		todo.GET("/list", h.parseListPage)
	}

	return router
}

func (h *Handler) refreshTokens(c *gin.Context) {
	refreshToken, exists := c.Get("refresh_token")
	if !exists {
		newResponseError(c, http.StatusBadRequest, "refresh token does not exist")
		return
	}
	id, exists := c.Get("id")
	if !exists {
		newResponseError(c, http.StatusBadRequest, "id does not exist")
		return
	}

	responseTokens, err := h.service.Authorization.RefreshAccessToken(id.(int), refreshToken.(string))
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.setAuthCookie(c, id.(int), responseTokens.RefreshToken)
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

	h.setAuthCookie(c, u.Id, responseTokens.RefreshToken)
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

	h.setAuthCookie(c, id, responseTokens.RefreshToken)
	c.JSON(http.StatusOK, *responseTokens)
}

func (h *Handler) setAuthCookie(c *gin.Context, id int, refreshToken string) {
	authCookie, err := h.store.Get(c.Request, "auth-cookie")
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	authCookie.Options = &sessions.Options{
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	authCookie.Values["id"] = id
	authCookie.Values["refresh_token"] = refreshToken

	if err = authCookie.Save(c.Request, c.Writer); err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) parseListPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
