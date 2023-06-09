package redis

import (
	"context"
	"fmt"

	cacheredis "github.com/joseclaudioads/url-shortener/internal/infra/cache"
	"github.com/joseclaudioads/url-shortener/internal/repositories/repository"
	"github.com/redis/go-redis/v9"
)

type UrlsCacheRedis struct{}

func NewUrlsCacheRedis() *UrlsCacheRedis {
	um := &UrlsCacheRedis{}
	return um
}

func (um UrlsCacheRedis) Save(s repository.ShortUrl) error {
	err := cacheredis.GetClient().Set(context.TODO(), s.Hash, s.OriginalUrl, 0).Err()
	if err != nil {
		panic(err)
	}

	return nil
}

func (um UrlsCacheRedis) Get(h string) (repository.ShortUrl, error) {
	result, err := cacheredis.GetClient().Get(context.TODO(), h).Result()

	if err == redis.Nil {
		fmt.Println("result not found")
	} else if err != nil {
		panic(err)
	}

	return repository.ShortUrl{
		OriginalUrl: result,
		Hash:        h,
	}, nil
}
