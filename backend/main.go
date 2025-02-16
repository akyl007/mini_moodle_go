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
	// Загрузка конфигурации
	if err := config.LoadConfig("backend/config/config.json"); err != nil {
		log.Fatal("Cannot load config:", err)
	}

	// База данных
	if err := db.Connect(); err != nil {
		log.Fatal("Cannot connect to database:", err)
	}
	defer db.DB.Close()

	// Настраиваем маршруты
	router := routes.SetupRouter()

	// Подключаем статическую раздачу из папки "frontend"
	fs := http.FileServer(http.Dir("frontend"))
	// Любой запрос, не начинающийся на /api, будет искать файл в папке "frontend"
	router.PathPrefix("/").Handler(http.StripPrefix("/", fs))

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

	// Запускаем сервер
	serverAddr := fmt.Sprintf(":%d", config.AppConfig.Server.Port)
	log.Printf("Server starting on http://localhost%s", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, handler))
}
