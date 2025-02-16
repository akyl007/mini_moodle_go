package handlers

import (
	"encoding/json"
	"log"
	"mini_moodle/backend/db"
	"mini_moodle/backend/models"
	"net/http"
)

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

func AssignTeacher(w http.ResponseWriter, r *http.Request) {
	var req struct {
		CourseID  int `json:"course_id"`
		TeacherID int `json:"teacher_id"`
	}
	log.Println("This line must work!")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}
	log.Printf("AssignTeacher request: course_id=%d, teacher_id=%d", req.CourseID, req.TeacherID)
	// Проверяем, что преподаватель с таким id существует и имеет роль teacher
	var teacherExists bool
	err := db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = $1 AND role = 'teacher')",
		req.TeacherID).Scan(&teacherExists)
	if err != nil {
		http.Error(w, "Ошибка проверки преподавателя", http.StatusInternalServerError)
		return
	}
	if !teacherExists {
		http.Error(w, "Преподаватель не найден", http.StatusBadRequest)
		return
	}
	log.Printf("AssignTeacher request: course_id=%d, teacher_id=%d", req.CourseID, req.TeacherID)
	// Проверяем, что курс существует
	var courseExists bool
	err = db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM courses WHERE id = $1)", req.CourseID).Scan(&courseExists)
	if err != nil {
		http.Error(w, "Ошибка проверки курса", http.StatusInternalServerError)
		return
	}
	if !courseExists {
		http.Error(w, "Курс не найден", http.StatusBadRequest)
		return
	}

	log.Printf("AssignTeacher request: course_id=%d, teacher_id=%d", req.CourseID, req.TeacherID)

	// Назначаем преподавателя на курс
	_, err = db.DB.Exec("UPDATE courses SET teacher_id = $1 WHERE id = $2",
		req.TeacherID, req.CourseID)
	if err != nil {
		http.Error(w, "Ошибка назначения преподавателя", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Преподаватель успешно назначен на курс"})
}
