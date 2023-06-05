package rest

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joseclaudioads/url-shortener/internal/http/handlers"
	"github.com/joseclaudioads/url-shortener/internal/services"
)

type ShortenerServer struct {
	services.ShortUrlService
}

func CreateServer(s ShortenerServer) *chi.Mux {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	h, err := handlers.NewUrlsHandler(s.ShortUrlService)

	if err != nil {
		log.Fatal("error creating urls handler")
		return nil
	}

	r.Post("/", h.CreateShortUrlHandler)
	r.Get("/{hash}", h.GetOriginalUrlHandler)
	return r
}
