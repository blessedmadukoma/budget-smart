package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewDBStorage(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("error initializing db:", err)
		return nil, err
	}

	return db, nil
}
