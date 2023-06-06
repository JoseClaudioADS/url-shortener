package services

import (
	"errors"

	"github.com/joseclaudioads/url-shortener/internal/repositories/repository"
	"github.com/joseclaudioads/url-shortener/internal/utils/idgenerator"
	"github.com/jxskiss/base62"
)

type ShortUrlService struct {
	repository.UrlRepository
}

func NewShortUrlService(u repository.UrlRepository) (*ShortUrlService, error) {
	if u == nil {
		return nil, errors.New("url repository not provided")
	}

	svc := &ShortUrlService{
		UrlRepository: u,
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

	h := e[0:7]

	s.UrlRepository.Save(repository.ShortUrl{
		OriginalUrl: o,
		Hash:        h,
	})
	return id, nil
}

func (s ShortUrlService) GetOriginalUrl(h string) (string, error) {

	u, error := s.UrlRepository.Get(h)

	if error != nil {
		return "", error
	}

	return u.OriginalUrl, nil
}
