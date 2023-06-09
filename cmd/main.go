package main

import (
	"log"
	"net/http"

	"github.com/joseclaudioads/url-shortener/internal/http/rest"
	"github.com/joseclaudioads/url-shortener/internal/repositories/caches/cache"
	"github.com/joseclaudioads/url-shortener/internal/repositories/caches/redis"
	"github.com/joseclaudioads/url-shortener/internal/repositories/mongo"
	"github.com/joseclaudioads/url-shortener/internal/repositories/postgres"
	"github.com/joseclaudioads/url-shortener/internal/repositories/repository"
	"github.com/joseclaudioads/url-shortener/internal/services"
	"github.com/joseclaudioads/url-shortener/internal/utils/environments"
)

func main() {

	var r repository.UrlRepository

	if environments.RepositoryType == "MONGO" {
		r = mongo.NewUrlsRepositoryMongo()
	} else {
		r = postgres.NewUrlRepositoryPostgres()
	}

	var c cache.UrlCache

	c = redis.NewUrlsCacheRedis()

	urlService, _ := services.NewShortUrlService(r, c)

	s := rest.CreateServer(rest.ShortenerServer{
		ShortUrlService: *urlService,
	})

	err := http.ListenAndServe(":3000", s)

	if err != nil {
		log.Fatal("Server was initialized with error", err)
	}
}
