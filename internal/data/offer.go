package data

import (
	"context"

	"github.com/osamikoyo/matesome/internal/data/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OfferStorage interface{
	Add(models.Offer) error
	Get(username string) ([]models.Offer, error)
}

type Offers struct{
	collection *mongo.Collection
	ctx context.Context
}

func (o *Offers) Add(offer models.Offer) error {
	_, err := o.collection.InsertOne(o.ctx, offer)
	return err
}

func (o *Offers) Get(username string) ([]models.Offer, error) {
	filter, err := o.collection.Find(o.ctx, bson.M{
		"author" : username, 
	})
	if err != nil{
		return nil, err
	}

	var offers []models.Offer

	if err := filter.All(o.ctx, &offers);err != nil{
		return nil, err
	}
	return offers, nil
}