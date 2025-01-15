package data

import (
	"github.com/osamikoyo/matesome/internal/data/models"
	"github.com/osamikoyo/matesome/pkg/loger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserStorage interface {
	Add(user models.User) error
	Login(email, password string) (string, error)
	Get(username string) (models.User, error)
}
type QuestStorage interface {
	Add(quest models.Quest) error
	GetByUserId(userid uint64) ([]models.Quest, error)
	GetByHashTags(hashtags models.HashTags) ([]models.Quest, error)
	Delete(quest models.Quest) error
}

func NewUserStorage() UserStorage {
	db, err := gorm.Open(sqlite.Open("storage/main.db"))
	if err != nil {
		loger.New().Error().Err(err)
	}

	db.AutoMigrate(&models.User{})

	return Users{db}
}
