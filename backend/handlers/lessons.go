package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"mini_moodle/backend/db"
	"mini_moodle/backend/models"
	"net/http"
	"strconv"
)

func DeleteLesson(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Lesson ID is required", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec("DELETE FROM lessons WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Failed to delete lesson", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Lesson with ID %s deleted successfully", id)
}

type AssignRequest struct {
	LessonID  int `json:"lesson_id"`
	TeacherID int `json:"teacher_id"`
}

func AssignTeacher(w http.ResponseWriter, r *http.Request) {
	var req AssignRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec("UPDATE lessons SET teacher_id = $1 WHERE id = $2", req.TeacherID, req.LessonID)
	if err != nil {
		http.Error(w, "Ошибка обновления урока", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Teacher assigned successfully"})
}
func GetLessons(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, name, description, teacher_id FROM lessons")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Database query error:", err)
		return
	}
	defer rows.Close()

	var lessons []models.Lesson
	for rows.Next() {
		var lesson models.Lesson
		var teacherID sql.NullInt64

		if err := rows.Scan(&lesson.ID, &lesson.Name, &lesson.Description, &teacherID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Row scan error:", err)
			return
		}

		if teacherID.Valid {
			id := int(teacherID.Int64)
			lesson.TeacherID = &id
		} else {
			lesson.TeacherID = nil
		}

		lessons = append(lessons, lesson)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lessons)
}
func GetTeachers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, username FROM users WHERE role = 'teacher'")
	if err != nil {
		http.Error(w, "Ошибка загрузки учителей", http.StatusInternalServerError)
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

func GetLesson(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Lesson ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Lesson ID", http.StatusBadRequest)
		return
	}

	var lesson models.Lesson
	var teacherID sql.NullInt64

	err = db.DB.QueryRow("SELECT id, name, description, teacher_id FROM lessons WHERE id = $1", id).
		Scan(&lesson.ID, &lesson.Name, &lesson.Description, &teacherID)

	if err != nil {
		http.Error(w, "Lesson not found", http.StatusNotFound)
		return
	}

	if teacherID.Valid {
		tID := int(teacherID.Int64)
		lesson.TeacherID = &tID
	} else {
		lesson.TeacherID = nil
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lesson)
}

type AssignStudentsRequest struct {
	LessonID   int   `json:"lesson_id"`
	StudentIDs []int `json:"student_ids"`
}

func AssignStudents(w http.ResponseWriter, r *http.Request) {
	var req AssignStudentsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for _, studentID := range req.StudentIDs {
		_, err := db.DB.Exec("INSERT INTO lesson_students (lesson_id, student_id) VALUES ($1, $2)", req.LessonID, studentID)
		if err != nil {
			http.Error(w, "Ошибка назначения студента", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Students assigned successfully"})
}
