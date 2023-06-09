package cacheredis

import (
	"context"

	"github.com/joseclaudioads/url-shortener/internal/utils/environments"
	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func GetClient() *redis.Client {

	if client != nil {
		return client
	}

	redisPassword := environments.RedisPassword
	redisUrl := environments.RedisUrl

	c := redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: redisPassword,
		DB:       0,
	})

	_, err := c.Ping(context.TODO()).Result()

	if err != nil {
		panic(err)
	}

	client = c

	return client
}
