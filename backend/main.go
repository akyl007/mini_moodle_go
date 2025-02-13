package main

import (
	"fmt"
	"log"
	"mini_moodle/backend/config"
	"mini_moodle/backend/db"
	"mini_moodle/backend/routes"
	"net/http"
)

func main() {
	// Загружаем конфигурацию
	if err := config.LoadConfig("backend/config/config.json"); err != nil {
		log.Fatal("Cannot load config:", err)
	}

	// Подключаемся к базе данных
	if err := db.Connect(); err != nil {
		log.Fatal("Cannot connect to database:", err)
	}
	defer func() {
		if err := db.DB.Close(); err != nil {
			log.Printf("Error closing DB: %v", err)
		}
	}()

	// Настраиваем маршрутизацию
	router := routes.SetupRouter()
	// Регистрируем обработчик статики
	fs := http.FileServer(http.Dir("frontend"))
	// Обратите внимание: этот обработчик срабатывает, если путь не обработан API-роутами.
	router.PathPrefix("/").Handler(fs)

	// Настраиваем CORS
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		router.ServeHTTP(w, r)
	})

	serverAddr := fmt.Sprintf(":%d", config.AppConfig.Server.Port)
	log.Printf("Server starting on http://localhost%s", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, handler))
}
