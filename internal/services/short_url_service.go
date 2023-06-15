package services

import (
	"errors"

	"github.com/joseclaudioads/url-shortener/internal/repositories/caches/cache"
	"github.com/joseclaudioads/url-shortener/internal/repositories/repository"
	"github.com/joseclaudioads/url-shortener/internal/utils/environments"
	"github.com/joseclaudioads/url-shortener/internal/utils/idgenerator"
	"github.com/jxskiss/base62"
)

type ShortUrlService struct {
	repository.UrlRepository
	cache.UrlCache
}

func NewShortUrlService(u repository.UrlRepository, c cache.UrlCache) (*ShortUrlService, error) {
	if u == nil {
		return nil, errors.New("url repository not provided")
	}

	svc := &ShortUrlService{
		UrlRepository: u,
		UrlCache:      c,
	}

	return svc, nil
}

func (s ShortUrlService) CreateShortUrl(o string) (string, error) {

	ig := idgenerator.IDGenerator{}

	id, err := ig.CreateID()

	if err != nil {
		return "", errors.New("Error generating id")
	}

	e := base62.EncodeToString([]byte(id))

	h := e[0:10]

	s.UrlRepository.Save(repository.ShortUrl{
		OriginalUrl: o,
		Hash:        h,
	})

	return h, nil
}

func (s ShortUrlService) GetOriginalUrl(h string) (string, error) {

	var u repository.ShortUrl
	var err error

	if environments.IsCacheEnable() {
		u, err = s.UrlCache.Get(h)
	}

	if !environments.IsCacheEnable() || err != nil || u.OriginalUrl == "" {
		u, err = s.UrlRepository.Get(h)

		if err != nil {
			return "", err
		}

		if u.OriginalUrl != "" && environments.IsCacheEnable() {
			go s.UrlCache.Save(u)
		}
	}

	return u.OriginalUrl, nil
}
