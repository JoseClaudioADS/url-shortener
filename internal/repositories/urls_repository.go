package repositories

import "fmt"

type ShortUrl struct {
	OriginalUrl string
	Hash        string
}

type UrlRepository interface {
	Save(s ShortUrl) error
}

type UrlRepositoryPostgres struct{}

func (up UrlRepositoryPostgres) Save(s ShortUrl) error {
	fmt.Println(s)
	return nil
}
