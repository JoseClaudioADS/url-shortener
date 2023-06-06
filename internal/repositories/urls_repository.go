package repositories

import "fmt"

type ShortUrl struct {
	OriginalUrl string
	Hash        string
}

type UrlRepository interface {
	Save(s ShortUrl) error
	Get(h string) (ShortUrl, error)
}

type UrlRepositoryPostgres struct{}

func (up UrlRepositoryPostgres) Save(s ShortUrl) error {
	fmt.Println(s)
	return nil
}

func (up UrlRepositoryPostgres) Get(h string) (ShortUrl, error) {
	s := ShortUrl{
		OriginalUrl: "https://google.com",
		Hash:        h,
	}
	fmt.Println(s)
	return s, nil
}
