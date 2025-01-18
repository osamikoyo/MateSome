package service

import (
	"github.com/osamikoyo/matesome/internal/data"
	"github.com/osamikoyo/matesome/internal/data/models"
)

type OfferValidator interface{
	Add(models.Offer) error
	Get(email string) ([]models.Offer, error)
}

type OfferService struct{
	Storage data.OfferStorage
}

func (o *OfferService) Add(offer models.Offer) error {
	return o.Storage.Add(offer)
}

func (o *OfferService) Get(username string) ([]models.Offer, error) {
	return o.Storage.Get(username)
}
