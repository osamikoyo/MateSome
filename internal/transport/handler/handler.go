package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/osamikoyo/matesome/internal/data"
	"github.com/osamikoyo/matesome/internal/service"
)

type Handler struct {
	User  UserHandler
	Quest QuestHandler
}

func New() Handler {
	return Handler{
		User: UserHandler{
			service.UserService{
				Storage: data.NewUserStorage(),
			},
		},
		Quest: QuestHandler{
			service.QuestService{
				QuestStorage: data.NewQuestStorage(),
			},
		},
	}
}

func (h Handler) RegisterRoutes(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	protect := e.Group("/quests", JWTMiddleware)
	protect.POST("/add", h.Quest.addQuestHandler)
	protect.POST("/delete", h.Quest.deleteQuestHandler)
	protect.GET("/self", h.Quest.getByUserIdHandler)

	e.POST("/search/hashtags", h.Quest.getByHashTagsHandler)

	e.POST("/register", h.User.registerHandler)
	e.POST("/login", h.User.loginHandler)
	e.GET("/profile/:username", h.User.getProfileHandler)
}
