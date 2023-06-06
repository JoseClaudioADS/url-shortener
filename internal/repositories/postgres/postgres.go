package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/joseclaudioads/url-shortener/internal/repositories/repository"
	"github.com/joseclaudioads/url-shortener/internal/utils/environments"

	_ "github.com/lib/pq"
)

type UrlRepositoryPostgres struct{}

func getUrlConnection() string {
	host := environments.DbHost
	port := environments.DbPort
	user := environments.DbUser
	password := environments.DbPassword
	dbname := environments.DbName

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

func openDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", getUrlConnection())
	if err != nil {
		return nil, err
	}

	maxConn := environments.GetDbMaxConnections()

	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxConn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func closeDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func (up UrlRepositoryPostgres) Save(s repository.ShortUrl) error {
	db, err := openDB()
	defer closeDB(db)
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

	return nil
}

func (up UrlRepositoryPostgres) Get(h string) (repository.ShortUrl, error) {
	s := repository.ShortUrl{
		Hash: h,
	}

	db, err := openDB()
	defer closeDB(db)
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

	return s, nil
}
