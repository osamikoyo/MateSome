package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/osamikoyo/matesome/internal/data/models"
	"github.com/osamikoyo/matesome/internal/service"
	"net/http"
)

type UserHandler struct {
	Service service.UserValidator
}

func (h UserHandler) registerHandler(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if err := h.Service.Register(user); err != nil {
		return err
	}

	return c.String(http.StatusCreated, "success!")
}

func (h UserHandler) loginHandler(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	token, err := h.Service.Login(email, password)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

func (h UserHandler) getProfileHandler(c echo.Context) error {
	username := c.QueryParam("username")
	user, err := h.Service.Get(username)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}
