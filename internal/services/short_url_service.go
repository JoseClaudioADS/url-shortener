package services

import (
	"errors"
	"fmt"

	"github.com/joseclaudioads/url-shortener/internal/repositories"
)

type ShortUrlService struct {
	repositories.UrlRepository
}

func NewShortUrlService(u repositories.UrlRepository) (*ShortUrlService, error) {
	if u == nil {
		return nil, errors.New("url repository not provided")
	}

	svc := &ShortUrlService{
		UrlRepository: u,
	}

	return svc, nil
}

func (s ShortUrlService) CreateShortUrl(o string) (string, error) {
	fmt.Printf("Original Url %s", o)

	h := "jdj49f"

	s.UrlRepository.Save(repositories.ShortUrl{
		OriginalUrl: o,
		Hash:        h,
	})
	return h, nil
}
