package handlers

import (
	"encoding/json"
	"mini_moodle/backend/db"
	"mini_moodle/backend/models"
	"net/http"
)

func GetStudents(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, username FROM users WHERE role = 'student'")
	if err != nil {
		http.Error(w, "Ошибка загрузки студентов", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var student models.Student
		if err := rows.Scan(&student.ID, &student.Username); err != nil {
			http.Error(w, "Ошибка обработки данных", http.StatusInternalServerError)
			return
		}
		students = append(students, student)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}
