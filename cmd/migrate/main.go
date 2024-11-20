package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func loadEnvVariables() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("Error loading .env file")
	}
	return nil
}

func main() {
	err := loadEnvVariables()
	if err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	defer db.Close()

	m, err := migrate.New(
		"file://migrations",
		connStr,
	)

	if err != nil {
		log.Fatalf("failed to create migrations instance: %v", err)
	}

	if len(os.Args) < 2 {
		log.Fatal("Please specify a migration action: up, down, or reset.")
	}

	action := os.Args[1]
	switch action {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("failed to apply migrations: %v", err)
		}
		fmt.Println("Migrations applied successfully.")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("failed to rollback migrations: %v", err)
		}
		fmt.Println("Migrations rolled back successfully.")
	case "reset":
		if err := m.Drop(); err != nil {
			log.Fatalf("failed to reset migrations: %v", err)
		}
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("failed to re-apply migrations: %v", err)
		}
		fmt.Println("Migrations reset successfully.")
	default:
		log.Fatalf("Unknown action: %s", action)
	}
}
