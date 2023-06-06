package postgres

import (
	"database/sql"
	"fmt"

	"github.com/joseclaudioads/url-shortener/internal/repositories/repository"
	"github.com/joseclaudioads/url-shortener/internal/utils/environments"

	_ "github.com/lib/pq"
)

type UrlRepositoryPostgres struct{}

var db *sql.DB

func getUrlConnection() string {
	host := environments.DbHost
	port := environments.DbPort
	user := environments.DbUser
	password := environments.DbPassword
	dbname := environments.DbName

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

func getDBConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", getUrlConnection())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("successful connection!")
	return db, nil
}

func (up UrlRepositoryPostgres) Save(s repository.ShortUrl) error {
	fmt.Println(s)

	db, err := getDBConnection()
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = db.Exec("INSERT INTO urls (hash, original_url) VALUES ($1, $2)", s.Hash, s.OriginalUrl)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("INSERT executed successfully!")

	defer db.Close()

	return nil
}

func (up UrlRepositoryPostgres) Get(h string) (repository.ShortUrl, error) {
	s := repository.ShortUrl{
		Hash: h,
	}

	db, err := getDBConnection()
	if err != nil {
		fmt.Println(err)
		return s, err
	}

	query := "SELECT original_url FROM urls WHERE hash = $1"
	row := db.QueryRow(query, h)

	err = row.Scan(&s.OriginalUrl)
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

	defer db.Close()

	return s, nil
}
