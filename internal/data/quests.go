package data

import (
	"context"
	"github.com/osamikoyo/matesome/internal/data/models"
	"github.com/osamikoyo/matesome/pkg/loger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

type Quests struct {
	*mongo.Client
}

func (q Quests) Add(quest models.Quest) error {
	//TODO implement me
	panic("implement me")
}

func (q Quests) GetByUserId(userid uint64) ([]models.Quest, error) {
	//TODO implement me
	panic("implement me")
}

func (q Quests) GetByHashTags(hashtags models.HashTags) ([]models.Quest, error) {
	//TODO implement me
	panic("implement me")
}

func (q Quests) Delete(quest models.Quest) error {
	//TODO implement me
	panic("implement me")
}

func NewQuestStorage() Quests {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		loger.New().Error().Err(err)
	}

	return Quests{client}
}
