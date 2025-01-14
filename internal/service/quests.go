package service

import (
	"github.com/osamikoyo/matesome/internal/data"
	"github.com/osamikoyo/matesome/internal/data/models"
)

type QuestsValidator interface {
	Add(quest models.Quest) error
	GetByUserId(userid uint64) ([]models.Quest, error)
	GetByHashTags(hashtags models.HashTags) ([]models.Quest, error)
	Delete(quest models.Quest) error
}

type QuestService struct {
	data.QuestStorage
}
