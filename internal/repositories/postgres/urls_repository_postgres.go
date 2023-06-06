package postgres

import (
	"database/sql"
	"fmt"

	"github.com/joseclaudioads/url-shortener/internal/infra/db"
	"github.com/joseclaudioads/url-shortener/internal/repositories/repository"

	_ "github.com/lib/pq"
)

type UrlRepositoryPostgres struct {
	db *sql.DB
}

func NewUrlRepositoryPostgres() *UrlRepositoryPostgres {
	db, err := db.OpenDB()

	if err != nil {
		panic(err)
	}

	urp := &UrlRepositoryPostgres{
		db: db,
	}

	return urp
}

func (up UrlRepositoryPostgres) Save(s repository.ShortUrl) error {

	_, err := up.db.Exec("INSERT INTO urls (hash, original_url) VALUES ($1, $2)", s.Hash, s.OriginalUrl)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("INSERT executed successfully!")

	return nil
}

func (up UrlRepositoryPostgres) Get(h string) (repository.ShortUrl, error) {
	s := repository.ShortUrl{
		Hash: h,
	}

	query := "SELECT original_url FROM urls WHERE hash = $1"
	row := up.db.QueryRow(query, h)

	err := row.Scan(&s.OriginalUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no results found")
		} else {
			fmt.Println(err)
			return s, err
		}
	} else {
		fmt.Println("results found:")
		fmt.Println(s.OriginalUrl)
	}

	fmt.Println(s)

	return s, nil
}
