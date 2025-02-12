package handlers

import (
	"encoding/json"
	"mini_moodle/backend/db"
	"mini_moodle/backend/models"
	"net/http"
)

// GetTeachers возвращает список всех преподавателей
func GetTeachers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, username FROM users WHERE role = 'teacher'")
	if err != nil {
		http.Error(w, "Ошибка загрузки преподавателей", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var teachers []models.Teacher
	for rows.Next() {
		var teacher models.Teacher
		if err := rows.Scan(&teacher.ID, &teacher.Username); err != nil {
			http.Error(w, "Ошибка обработки данных", http.StatusInternalServerError)
			return
		}
		teachers = append(teachers, teacher)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teachers)
}

// AssignTeacher назначает преподавателя на урок
type AssignTeacherRequest struct {
	LessonID  int `json:"lesson_id"`
	TeacherID int `json:"teacher_id"`
}

func AssignTeacher(w http.ResponseWriter, r *http.Request) {
	var req AssignTeacherRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec("UPDATE lessons SET teacher_id = $1 WHERE id = $2",
		req.TeacherID, req.LessonID)
	if err != nil {
		http.Error(w, "Ошибка назначения преподавателя", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Преподаватель успешно назначен"})
}
