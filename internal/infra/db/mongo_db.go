package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joseclaudioads/url-shortener/internal/utils/environments"
)

var client *mongo.Client

func GetClient() *mongo.Client {

	if client != nil {
		return client
	}

	mongo_url := environments.MongoUrl

	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongo_url))
	if err != nil {
		panic(err)
	}

	client = c

	return client
}
