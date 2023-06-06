package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/joseclaudioads/url-shortener/internal/services"
)

type createShortUrlInput struct {
	OriginalUrl string
}

type urlsHandler struct {
	services.ShortUrlService
}

func NewUrlsHandler(s services.ShortUrlService) (*urlsHandler, error) {
	h := &urlsHandler{
		ShortUrlService: s,
	}

	return h, nil
}

func (u urlsHandler) CreateShortUrlHandler(w http.ResponseWriter, r *http.Request) {

	var i createShortUrlInput

	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u.ShortUrlService.CreateShortUrl(i.OriginalUrl)
	w.WriteHeader(201)
	w.Write([]byte("Short URL Created"))
}

func (u urlsHandler) GetOriginalUrlHandler(w http.ResponseWriter, r *http.Request) {
	hashParam := chi.URLParam(r, "hash")
	if strings.TrimSpace(hashParam) == "" {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("you must give a valid url hash")))
		return
	}

	o, error := u.ShortUrlService.GetOriginalUrl(hashParam)

	if error != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("something wrong happened")))
		return
	}

	w.Header().Set("Location", o)

	w.WriteHeader(http.StatusMovedPermanently)
}
