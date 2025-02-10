package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go" // JWT для токенов
	"golang.org/x/crypto/bcrypt"  // bcrypt для хеширования паролей
	"mini_moodle/backend/db"      // пакет для работы с базой
	"mini_moodle/backend/models"  // модели (например, User)
)

// Секрет для JWT — лучше вынести в конфигурацию или переменные окружения
var jwtSecret = []byte("your-secret-key")

// RegisterUser принимает JSON с username, password и role, хеширует пароль и вставляет пользователя в базу.
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	// Декодируем JSON из тела запроса
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Если пароль пустой — возвращаем ошибку
	if user.Password == "" {
		http.Error(w, "Password is required", http.StatusBadRequest)
		return
	}

	// Хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error generating hash: %v", err)
		http.Error(w, "Password hashing failed", http.StatusInternalServerError)
		return
	}

	// Вставляем пользователя в базу; убедись, что таблица users имеет достаточно места для хэша (например, тип TEXT или VARCHAR(100))
	_, err = db.DB.Exec(
		"INSERT INTO users (username, password, role) VALUES ($1, $2, $3)",
		user.Username, string(hashedPassword), user.Role,
	)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// LoginUser принимает JSON с username и password, сравнивает введённый пароль с хешем из базы и генерирует JWT.
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	// Декодируем JSON с учетными данными
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	log.Printf("Login attempt: username=%s, password=%s", creds.Username, creds.Password)

	var user models.User
	// Извлекаем пользователя из базы. Если таблица называется "users", можно не оборачивать имя в кавычки,
	// но если у тебя PostgreSQL чувствителен к регистру — убедись, что имя таблицы указано правильно.
	err := db.DB.QueryRow(
		"SELECT id, username, password, role FROM users WHERE username = $1",
		creds.Username,
	).Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		log.Printf("Query error: %v", err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Выводим для отладки: длина и само значение хэша
	log.Printf("Length of stored hash: %d", len(user.Password))
	log.Printf("Stored hash for user %s: %s", user.Username, user.Password)

	// Здесь предполагаем, что значение user.Password не содержит лишних символов,
	// но если нужно — можно использовать strings.TrimSpace для очистки.
	cleanHash := user.Password // либо strings.TrimSpace(user.Password)
	// Сравниваем полученный хэш с введённым паролем
	if err := bcrypt.CompareHashAndPassword([]byte(cleanHash), []byte(creds.Password)); err != nil {
		log.Printf("Password mismatch: %v", err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Если пароль совпадает, генерируем JWT-токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		http.Error(w, "Token generation failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
