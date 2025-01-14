package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/osamikoyo/matesome/internal/data/models"
	"github.com/osamikoyo/matesome/internal/service"
	"net/http"
)

type QuestHandler struct {
	service.QuestsValidator
}

func (h QuestHandler) addQuestHandler(c echo.Context) error {
	var quest models.Quest
	if err := c.Bind(&quest); err != nil {
		return err
	}

	if err := h.Add(quest); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusCreated, "success!")
}

func (h QuestHandler) getByHashTagsHandler(c echo.Context) error {
	var tags models.HashTags

	if err := c.Bind(&tags); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	quests, err := h.GetByHashTags(tags)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, quests)
}

func (h QuestHandler) getByUserIdHandler(c echo.Context) error {
	email := c.Get("email").(string)
	quests, err := h.GetByUserId(email)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, quests)
}

func (h QuestHandler) deleteQuestHandler(c echo.Context) error {
	var quest models.Quest
	if err := c.Bind(&quest); err != nil {
		return err
	}

	if err := h.Delete(quest); err != nil {
		return err
	}

	return c.String(http.StatusOK, "success!")
}
