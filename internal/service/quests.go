package service

import (
	"github.com/osamikoyo/matesome/internal/data"
	"github.com/osamikoyo/matesome/internal/data/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type QuestService struct {
	data.QuestStorage
}

type QuestsValidator interface {
	Add(quest models.Quest) error
	GetByUserId(email string) ([]models.Quest, error)
	GetByHashTags(hashtags models.HashTags) ([]models.Quest, error)
	Delete(quest models.Quest) error
}

func (q QuestService) Add(quest models.Quest) error {
	return q.QuestStorage.Add(quest)
}

func (q QuestService) GetByUserId(email string) ([]models.Quest, error) {

	db, err := gorm.Open(sqlite.Open("storage/main.db"))
	if err != nil {
		return nil, err
	}
	var user models.User
	if err = db.Where(&models.User{Email: email}).Find(&user).Error; err != nil {
		return nil, err
	}

	return q.QuestStorage.GetByUserId(user.ID)
}
func (q QuestService) GetByHashTags(hashtags models.HashTags) ([]models.Quest, error) {
	return q.QuestStorage.GetByHashTags(hashtags)
}

func (q QuestService) Delete(quest models.Quest) error {
	return q.QuestStorage.Delete(quest)
}
