package data

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/osamikoyo/matesome/internal/data/models"
	"gorm.io/gorm"
	"os"
)

type Users struct {
	*gorm.DB
}

func (u Users) Add(user models.User) error {
	return u.Create(&user).Error
}

func (u Users) Login(email, password string) (string, error) {
	var user models.User
	if err := u.Where(
		&models.User{
			Email:    email,
			Password: password,
		},
	).Find(&user).Error; err != nil {
		return "", err
	}

	claims := &models.Claims{Email: email}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(os.Getenv("JWT_KEY"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (u Users) Get(email string) (models.User, error) {
	var user models.User
	if err := u.Where(
		&models.User{
			Email: email,
		}).Find(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}
