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
	data.UserStorage
}
