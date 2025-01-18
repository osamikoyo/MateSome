package service

import (
	"github.com/osamikoyo/matesome/internal/data"
	"github.com/osamikoyo/matesome/internal/data/models"
	"golang.org/x/crypto/bcrypt"
)


type UserValidator interface {
	Register(user models.User) error
	Login(email, password string) (string, error)
	Get(email string) (models.User, error)
}
type UserService struct {
	Storage data.UserStorage
}

func (u UserService) Register(user models.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)

	return u.Storage.Add(user)
}

func (u UserService) Login(email, password string) (string, error) {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return u.Storage.Login(email, string(hashpassword))
}

func (u UserService) Get(email string) (models.User, error) {
	return u.Storage.Get(email)
}
