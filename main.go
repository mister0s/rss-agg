package main

import (
	"context"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/mister0s/rss-agg/internal/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is empty!!!")
	}

	dbUrl := os.Getenv("DB_URL")
	if port == "" {
		log.Fatal("DB_URL is empty!!!")
	}

	dbpool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	defer dbpool.Close()

	apiServer := NewApiServer(port, database.New(dbpool))
	apiServer.Run()
}
