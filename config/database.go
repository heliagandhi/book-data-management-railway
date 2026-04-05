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

	// Initialize tables if they don't exist
	InitTables()
}

func InitTables() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			created_by VARCHAR(50),
			modified_at TIMESTAMP DEFAULT NULL,
			modified_by VARCHAR(50)
		);`,
		`CREATE TABLE IF NOT EXISTS categories (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			created_by VARCHAR(50),
			modified_at TIMESTAMP DEFAULT NULL,
			modified_by VARCHAR(50)
		);`,
		`CREATE TABLE IF NOT EXISTS books (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			category_id INT NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
			description TEXT,
			image_url VARCHAR(255),
			release_year INT CHECK (release_year > 0),
			price DECIMAL(12,2) NOT NULL DEFAULT 0,
			total_page INT CHECK (total_page > 0),
			thickness VARCHAR(20),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			created_by VARCHAR(50),
			modified_at TIMESTAMP NULL,
			modified_by VARCHAR(50)
		);`,
		`CREATE TABLE IF NOT EXISTS user_sessions (
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			token TEXT NOT NULL,
			expired_at TIMESTAMP NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,
	}

	for _, q := range queries {
		_, err := DB.Exec(q)
		if err != nil {
			log.Fatal("Error creating table: ", err)
		}
	}

	log.Println("All tables ensured (created if not exists)")
}