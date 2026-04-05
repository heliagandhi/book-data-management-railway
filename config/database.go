package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	connStr := os.Getenv("DATABASE_URL")

	if connStr == "" {
		log.Println("Using local database")
		connStr = "host=localhost port=5432 user=postgres password=admin dbname=book_data_management sslmode=disable"
	}

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("DB connection error:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("DB not reachable:", err)
	}

	log.Println("Database connected")
}