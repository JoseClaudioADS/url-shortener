package cache

import "github.com/joseclaudioads/url-shortener/internal/repositories/repository"

type UrlCache interface {
	Save(s repository.ShortUrl) error
	Get(h string) (repository.ShortUrl, error)
}
