package repository

type ShortUrl struct {
	OriginalUrl string
	Hash        string
}

type UrlRepository interface {
	Save(s ShortUrl) error
	Get(h string) (ShortUrl, error)
}
