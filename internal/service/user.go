package service

import (
	"github.com/osamikoyo/matesome/internal/data"
	"github.com/osamikoyo/matesome/internal/data/models"
)

type UserValidator interface {
	Register(user models.User) error
	Login(email, password string) (string, error)
}
type UserService struct {
	Storage data.UserStorage
}

func (u UserService) Register(user models.User) error {
	return u.Storage.Add(user)
}

func (u UserService) Login(email, password string) (string, error) {
	return u.Storage.Login(email, password)
}
