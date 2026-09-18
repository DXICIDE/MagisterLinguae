package main

import (
	"database/sql"
	"log"
	"os"

	magisterlinguae "github.com/DXICIDE/MagisterLinguae"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // registers the "postgres" driver
	"github.com/pressly/goose/v3"
)

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("couldn't connect to the database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("couldn't ping the database: %v", err)
	}

	goose.SetDialect("postgres")

	goose.SetBaseFS(magisterlinguae.MigrationsFS)
	if err := goose.Up(db, "sql/schema"); err != nil {
		log.Fatalf("couldn't apply migrations: %v", err)
	}
}
