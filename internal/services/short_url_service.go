package services

import (
	"errors"
	"fmt"

	"github.com/joseclaudioads/url-shortener/internal/repositories/repository"
	"github.com/joseclaudioads/url-shortener/internal/utils/idgenerator"
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
	fmt.Println("Original Url %s", o)

	ig := idgenerator.IDGenerator{}

	id, err := ig.CreateID()

	if err != nil {
		return "", errors.New("Error generating id")
	}

	s.UrlRepository.Save(repository.ShortUrl{
		OriginalUrl: o,
		Hash:        id,
	})
	return id, nil
}

func (s ShortUrlService) GetOriginalUrl(h string) (string, error) {
	fmt.Println("Hash Url %s", h)

	u, error := s.UrlRepository.Get(h)

	if error != nil {
		return "", error
	}

	return u.OriginalUrl, nil
}
