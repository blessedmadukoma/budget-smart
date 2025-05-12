package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/blessedmadukoma/budgetsmart/engine/cmd/api"
	"github.com/blessedmadukoma/budgetsmart/engine/config"
	"github.com/blessedmadukoma/budgetsmart/engine/db"
)

func main() {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBAddress,
		config.Envs.DBName,
	)
	db, err := db.NewDBStorage(connStr)

	if err != nil {
		log.Fatal("error initializing db:", err)
	}

	initStorage(db)

	server := api.NewAPIServer(config.Envs.Port, db)

	if err := server.Run(); err != nil {
		log.Fatal("error running server:", err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB Successfully connected!")
}
