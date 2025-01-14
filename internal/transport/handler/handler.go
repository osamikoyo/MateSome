package handler

import (
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
				data.NewQuestStorage(),
			},
		},
	}
}
