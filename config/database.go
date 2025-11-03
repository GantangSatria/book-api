package config

import (
	"database/sql"
	//"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// user := os.Getenv("DB_USER")
	// pass := os.Getenv("DB_PASSWORD")
	// name := os.Getenv("DB_NAME")
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("No .env file found")
	}

	dsn := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("db open error: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("db ping error: %v", err)
	}

	return db
}