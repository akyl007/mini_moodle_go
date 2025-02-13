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

func TestCreateFeedback(t *testing.T) {
	// Тест отправки отзыва студентом
	t.Run("Create Feedback as Student", func(t *testing.T) {
		feedback := models.FeedbackRequest{
			CourseID:  1,
			TeacherID: 1,
			Comment:   "Great course!",
		}
		body, _ := json.Marshal(feedback)
		req := httptest.NewRequest("POST", "/api/feedback", bytes.NewBuffer(body))
		
		// Создаем тестовый JWT токен для студента
		token, _ := utils.GenerateToken(3, "teststudent", "student")
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		
		w := httptest.NewRecorder()

		handlers.CreateFeedback(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
		}
	})

	// Тест отправки отзыва преподавателем (должно быть запрещено)
	t.Run("Create Feedback as Teacher", func(t *testing.T) {
		feedback := models.FeedbackRequest{
			CourseID:  1,
			TeacherID: 1,
			Comment:   "Test feedback",
		}
		body, _ := json.Marshal(feedback)
		req := httptest.NewRequest("POST", "/api/feedback", bytes.NewBuffer(body))
		
		// Создаем тестовый JWT токен для преподавателя
		token, _ := utils.GenerateToken(4, "testteacher", "teacher")
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		
		w := httptest.NewRecorder()

		handlers.CreateFeedback(w, req)

		if w.Code != http.StatusForbidden {
			t.Errorf("Expected status code %d, got %d", http.StatusForbidden, w.Code)
		}
	})
} 