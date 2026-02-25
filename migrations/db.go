package migrations

import (
	"database/sql"
	"embed"
	"log"
)

// Встраиваем все SQL-файлы из текущей папки
//
//go:embed *.sql
var files embed.FS

func RunMigrations(db *sql.DB) {
	sqlBytes, err := files.ReadFile("001_init.sql") // имя файла должно совпадать
	if err != nil {
		log.Fatalf("Failed to read migration: %v", err)
	}
	if _, err := db.Exec(string(sqlBytes)); err != nil {
		log.Fatalf("Failed to run migration: %v", err)
	}
	log.Println("Migrations ran successfully!")
}
