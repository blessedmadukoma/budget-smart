package main

import (
	"fmt"
	"log"
	"os"

	"github.com/blessedmadukoma/budgetsmart/engine/config"
	"github.com/blessedmadukoma/budgetsmart/engine/db"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Migration direction (up/down) is required")
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBAddress,
		config.Envs.DBName,
	)
	_, err := db.NewDBStorage(connStr)
	if err != nil {
		log.Fatal("error initializing db:", err)
	}

	m, err := migrate.New(
		"file://db/migrations",
		connStr)

	if err != nil {
		log.Fatalf("Migration initialization error: %v", err)
	}

	direction := os.Args[1]

	if direction == "up" {
		log.Println("Running migrations up...")
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Failed to apply migrations: %v", err)
		}
	} else if direction == "down" {
		log.Println("Running migrations down...")
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Failed to roll back migrations: %v", err)
		}
	} else {
		log.Fatalf("Invalid direction '%s'. Use 'up' or 'down'", direction)
	}

	log.Printf("Migration '%s' completed successfully", direction)
}
