package tests

import (
	"bytes"
	"encoding/json"
	"mini_moodle/backend/handlers"
	"mini_moodle/backend/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	// Тест успешного входа
	t.Run("Successful Login", func(t *testing.T) {
		loginReq := handlers.LoginRequest{
			Username: "testuser",
			Password: "testpass",
		}
		body, _ := json.Marshal(loginReq)
		req := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handlers.Login(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}

		var response map[string]string
		json.NewDecoder(w.Body).Decode(&response)
		if _, exists := response["token"]; !exists {
			t.Error("Expected token in response")
		}
	})

	// Тест неверных учетных данных
	t.Run("Invalid Credentials", func(t *testing.T) {
		loginReq := handlers.LoginRequest{
			Username: "wronguser",
			Password: "wrongpass",
		}
		body, _ := json.Marshal(loginReq)
		req := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handlers.Login(w, req)

		if w.Code != http.StatusUnauthorized {
			t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, w.Code)
		}
	})
} 