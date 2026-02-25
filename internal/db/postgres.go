package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitPostgres() *sql.DB {
	// Загружаем .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using OS environment or defaults")
	}

	getEnv := func(key, fallback string) string {
		if value, ok := os.LookupEnv(key); ok {
			return value
		}
		return fallback
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "127.0.0.1"), // локальный запуск
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "marketing_user"),
		getEnv("DB_PASS", "pass"),
		getEnv("DB_NAME", "prices"),
		getEnv("DB_SSLMODE", "disable"),
	)

	var db *sql.DB
	var err error

	// Попытки подключения с таймаутом
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Println("Failed to open database, retrying...", err)
		} else if pingErr := db.Ping(); pingErr == nil {
			log.Println("Connected to Postgres successfully!")
			return db
		} else {
			log.Println("Waiting for Postgres to be ready...", pingErr)
		}
		time.Sleep(2 * time.Second)
	}

	log.Fatal("Could not connect to Postgres after multiple attempts:", err)
	return nil
}
