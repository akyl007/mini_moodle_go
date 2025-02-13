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

func TestCreateCourse(t *testing.T) {
	// Тест создания курса преподавателем
	t.Run("Create Course as Teacher", func(t *testing.T) {
		course := models.Course{
			Name:        "Test Course",
			Description: "Test Description",
		}
		body, _ := json.Marshal(course)
		req := httptest.NewRequest("POST", "/api/course", bytes.NewBuffer(body))
		
		// Создаем тестовый JWT токен для преподавателя
		token, _ := utils.GenerateToken(1, "testteacher", "teacher")
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		
		w := httptest.NewRecorder()

		handlers.CreateCourse(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
		}

		var response models.Course
		json.NewDecoder(w.Body).Decode(&response)
		if response.Name != course.Name {
			t.Errorf("Expected course name %s, got %s", course.Name, response.Name)
		}
	})

	// Тест создания курса студентом (должно быть запрещено)
	t.Run("Create Course as Student", func(t *testing.T) {
		course := models.Course{
			Name:        "Student Course",
			Description: "Test Description",
		}
		body, _ := json.Marshal(course)
		req := httptest.NewRequest("POST", "/api/course", bytes.NewBuffer(body))
		
		// Создаем тестовый JWT токен для студента
		token, _ := utils.GenerateToken(2, "teststudent", "student")
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		
		w := httptest.NewRecorder()

		handlers.CreateCourse(w, req)

		if w.Code != http.StatusForbidden {
			t.Errorf("Expected status code %d, got %d", http.StatusForbidden, w.Code)
		}
	})
} 