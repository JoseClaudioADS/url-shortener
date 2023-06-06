package mongo

import (
	"context"

	"github.com/joseclaudioads/url-shortener/internal/infra/db"
	"github.com/joseclaudioads/url-shortener/internal/repositories/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type UrlsRepositoryMongo struct{}

func NewUrlsRepositoryMongo() *UrlsRepositoryMongo {
	um := &UrlsRepositoryMongo{}
	return um
}

func (um UrlsRepositoryMongo) Save(s repository.ShortUrl) error {
	uc := db.GetClient().Database("urlshortener").Collection("urls")

	u := bson.D{{Key: "original_url", Value: s.OriginalUrl}, {Key: "hash", Value: s.Hash}}

	_, err := uc.InsertOne(context.TODO(), u)

	if err != nil {
		panic(err)
	}

	// result.InsertedID

	return nil
}

func (um UrlsRepositoryMongo) Get(h string) (repository.ShortUrl, error) {

	return repository.ShortUrl{}, nil
}
