package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joseclaudioads/url-shortener/internal/utils/environments"
)

func GetClient() *mongo.Client {

	mongo_url := environments.MongoUrl

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongo_url))
	if err != nil {
		panic(err)
	}

	return client
}
