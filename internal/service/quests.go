package service

import (
	"github.com/osamikoyo/matesome/internal/data"
	"github.com/osamikoyo/matesome/internal/data/models"
)

type QuestService struct {
	data.QuestStorage
}

type QuestsValidator interface {
	Add(quest models.Quest) error
	GetByUserId(userid uint64) ([]models.Quest, error)
	GetByHashTags(hashtags models.HashTags) ([]models.Quest, error)
	Delete(quest models.Quest) error
}

func (q QuestService) Add(quest models.Quest) error {
	return q.QuestStorage.Add(quest)
}

func (q QuestService) GetByUserId(userid uint64) ([]models.Quest, error) {
	return q.QuestStorage.GetByUserId(userid)
}
func (q QuestService) GetByHashTags(hashtags models.HashTags) ([]models.Quest, error) {
	return q.QuestStorage.GetByHashTags(hashtags)
}

func (q QuestService) Delete(quest models.Quest) error {
	return q.QuestStorage.Delete(quest)
}
