package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/blessedmadukoma/budgetsmart/engine/cmd/api"
	"github.com/blessedmadukoma/budgetsmart/engine/config"
	"github.com/blessedmadukoma/budgetsmart/engine/db"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/log"
)

func main() {
	logger := log.NewLogger(os.Stdout)
	logger.SetPrefix("main")

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBAddress,
		config.Envs.DBName,
	)
	db, err := db.NewDBStorage(connStr)

	if err != nil {
		logger.WithError(err).Fatal("error initializing db")
	}

	initStorage(db, logger)

	server := api.NewAPIServer(config.Envs.Port, db, logger)

	if err := server.Run(); err != nil {
		logger.WithError(err).Fatal("error running server")
	}
}

func initStorage(db *sql.DB, logger *log.Logger) {
	err := db.Ping()

	if err != nil {
		logger.WithError(err).Fatal("failed to ping database")
	}

	logger.Info("DB Successfully connected!")
}
