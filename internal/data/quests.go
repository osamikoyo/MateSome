package data

import (
	"context"
	"fmt"
	"github.com/osamikoyo/matesome/internal/data/models"
	"github.com/osamikoyo/matesome/pkg/loger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

type Quests struct {
	ctx context.Context
	*mongo.Collection
}

func (q Quests) Add(quest models.Quest) error {
	_, err := q.InsertOne(q.ctx, quest)
	return err
}

func (q Quests) GetByUserId(userid uint64) ([]models.Quest, error) {
	filter, err := q.Find(q.ctx, bson.M{
		"user_id": fmt.Sprintf("%d", userid),
	})

	if err != nil {
		return nil, err
	}

	var quests []models.Quest

	if err = filter.All(q.ctx, &quests); err != nil {
		return nil, err
	}
	return quests, nil
}

func (q Quests) GetByHashTags(hashtags models.HashTags) ([]models.Quest, error) {
	filter, err := q.Find(q.ctx, bson.M{
		"hash_tag": bson.A{
			hashtags.Tags,
		},
	})

	if err != nil {
		return nil, err
	}

	var quests []models.Quest

	if err = filter.All(q.ctx, &quests); err != nil {
		return nil, err
	}
	return quests, nil
}

func (q Quests) Delete(quest models.Quest) error {
	_, err := q.DeleteOne(q.ctx, quest)
	return err
}

func NewQuestStorage() Quests {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		loger.New().Error().Err(err)
	}

	collection := client.Database("main").Collection("quests")

	return Quests{ctx: ctx, Collection: collection}
}
