package mongo

import (
	"context"
	"fmt"

	"github.com/joseclaudioads/url-shortener/internal/infra/db"
	"github.com/joseclaudioads/url-shortener/internal/repositories/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UrlsRepositoryMongo struct{}

type ShortUrl struct {
	OriginalUrl string `bson:"original_url"`
	Hash        string `bson:"hash"`
}

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

	return nil
}

func (um UrlsRepositoryMongo) Get(h string) (repository.ShortUrl, error) {
	uc := db.GetClient().Database("urlshortener").Collection("urls")

	filter := bson.D{{Key: "hash", Value: h}}

	var result ShortUrl
	err := uc.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("result not found: " + h)
		} else {
			panic(err)
		}
	}

	return repository.ShortUrl{
		OriginalUrl: result.OriginalUrl,
		Hash:        result.Hash,
	}, nil
}
