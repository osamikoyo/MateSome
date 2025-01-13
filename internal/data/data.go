package data

import "github.com/osamikoyo/matesome/internal/data/models"

type UserStorage interface {
	Add(user models.User) error
	Login(email, password string) (bool, error)
}
type QuestStorage interface {
	Add(quest models.Quest) error
	GetByUserId(userid uint64) ([]models.Quest, error)
	GetByHashTags(hashtags models.HashTags) ([]models.Quest, error)
	Delete(quest models.Quest) error
}
