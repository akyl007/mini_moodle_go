package main

import (
	"fmt"
	"log"
	"mini_moodle/backend/config"
	"mini_moodle/backend/db"
	"mini_moodle/backend/routes"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	cfg := config.Load()
	db.Connect(cfg)

	// Выводим рабочую директорию для отладки
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Working directory: %s", wd)

	// Собираем абсолютный путь к сборке React
	staticPath := filepath.Join(wd, "frontend", "react-frontend", "build")
	log.Printf("Serving static files from: %s", staticPath)

	router := routes.SetupRouter(staticPath)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
