package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"mini_moodle/backend/config"
)

var DB *sql.DB

func Connect(cfg *config.Config) {
	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBHost, cfg.DBPort,
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Database is not responding:", err)
	}

	fmt.Println("Connected to the database")
}
