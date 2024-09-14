package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mister0s/rss-agg/internal/database"

	_ "github.com/lib/pq"
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

	conn, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatal("Failed to Open Connection!")
	}

	if conn.Ping() != nil {
		log.Fatal("Failed to Ping DB!")
	}

	apiServer := NewApiServer(port, database.New(conn))
	apiServer.Run()
}
