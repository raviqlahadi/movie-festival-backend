package db

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
)

func RunMigrations() {
	m, err := migrate.New(
		"file://migrations",
		"postgress://user:password@localhost:5432/moviefestival?sslmode=disable",
	)

	if err != nil {
		log.Fatalf("Failed to initialize migrations: %v", err)
	}
}
