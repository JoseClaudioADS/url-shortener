package main

import (
	"log"
	"net/http"

	"github.com/joseclaudioads/url-shortener/internal/http/rest"
	"github.com/joseclaudioads/url-shortener/internal/repositories"
	"github.com/joseclaudioads/url-shortener/internal/services"
)

func main() {

	urlService, _ := services.NewShortUrlService(repositories.UrlRepositoryPostgres{})

	s := rest.CreateServer(rest.ShortenerServer{
		ShortUrlService: *urlService,
	})

	err := http.ListenAndServe(":3000", s)

	if err != nil {
		log.Fatal("Server was initialized with error", err)
	}
}
