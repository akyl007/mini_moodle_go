package tests

import (
    "bytes"
    "encoding/json"
    "mini_moodle/backend/handlers"
    "mini_moodle/backend/models"
    "mini_moodle/backend/utils"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestForumMessages(t *testing.T) {
    // Тест создания сообщения
    t.Run("Create Forum Message", func(t *testing.T) {
        message := models.CreateForumMessageRequest{
            Message: "Test forum message",
        }
        body, _ := json.Marshal(message)
        req := httptest.NewRequest("POST", "/api/forum/message", bytes.NewBuffer(body))
        
        token, _ := utils.GenerateToken(1, "testuser", "student")
        req.Header.Set("Authorization", "Bearer "+token)
        req.Header.Set("Content-Type", "application/json")
        
        w := httptest.NewRecorder()
        handlers.CreateForumMessage(w, req)

        if w.Code != http.StatusCreated {
            t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
        }
    })

    // Тест получения сообщений
    t.Run("Get Forum Messages", func(t *testing.T) {
        req := httptest.NewRequest("GET", "/api/forum/messages", nil)
        token, _ := utils.GenerateToken(1, "testuser", "student")
        req.Header.Set("Authorization", "Bearer "+token)
        
        w := httptest.NewRecorder()
        handlers.GetForumMessages(w, req)

        if w.Code != http.StatusOK {
            t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
        }

        var messages []models.ForumMessage
        json.NewDecoder(w.Body).Decode(&messages)
        if len(messages) == 0 {
            t.Error("Expected messages to be returned")
        }
    })
} 