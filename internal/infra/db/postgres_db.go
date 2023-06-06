package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/joseclaudioads/url-shortener/internal/utils/environments"
)

func getUrlConnection() string {
	host := environments.DbHost
	port := environments.DbPort
	user := environments.DbUser
	password := environments.DbPassword
	dbname := environments.DbName

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

func OpenDB() (*sql.DB, error) {
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
