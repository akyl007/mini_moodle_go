package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"mini_moodle/backend/db"
	"mini_moodle/backend/models"
	"net/http"
	"strconv"
)

// DeleteLesson удаляет урок по его ID
func DeleteLesson(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID урока обязателен", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID урока", http.StatusBadRequest)
		return
	}

	// Удаляем связи с таблицей lesson_students (если она есть)
	_, err = db.DB.Exec("DELETE FROM lesson_students WHERE lesson_id = $1", id)
	if err != nil {
		log.Printf("Ошибка удаления связей урока (lesson_students): %v", err)
		http.Error(w, "Ошибка удаления связей урока", http.StatusInternalServerError)
		return
	}

	// Удаляем сам урок
	result, err := db.DB.Exec("DELETE FROM lessons WHERE id = $1", id)
	if err != nil {
		log.Printf("Ошибка удаления урока: %v", err)
		http.Error(w, "Ошибка удаления урока", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Ошибка получения результата при удалении урока: %v", err)
		http.Error(w, "Ошибка получения результата", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Урок не найден", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Урок успешно удален"})
}

// GetLessons возвращает список всех уроков
func GetLessons(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(`SELECT id, name, description, teacher_id FROM lessons ORDER BY id`)
	if err != nil {
		log.Printf("Error querying lessons: %v", err)
		http.Error(w, "Error loading lessons: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var lessons []models.Lesson
	for rows.Next() {
		var lesson models.Lesson
		err := rows.Scan(
			&lesson.ID,
			&lesson.Name,
			&lesson.Description,
			&lesson.TeacherID,
		)
		if err != nil {
			log.Printf("Error scanning lesson: %v", err)
			http.Error(w, "Error processing lesson data: "+err.Error(), http.StatusInternalServerError)
			return
		}
		lessons = append(lessons, lesson)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating lessons: %v", err)
		http.Error(w, "Error processing lessons: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lessons)
}

// GetLesson возвращает информацию об одном уроке по ID
func GetLesson(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID урока обязателен", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID урока", http.StatusBadRequest)
		return
	}

	var lesson models.Lesson
	var teacherID sql.NullInt64
	var teacherUsername sql.NullString

	err = db.DB.QueryRow(`
        SELECT 
            l.id, 
            l.name, 
            l.description, 
            l.teacher_id,
            u.username as teacher_username
        FROM lessons l
        LEFT JOIN users u ON l.teacher_id = u.id AND u.role = 'teacher'
        WHERE l.id = $1
    `, id).Scan(
		&lesson.ID,
		&lesson.Name,
		&lesson.Description,
		&teacherID,
		&teacherUsername,
	)

	if err == sql.ErrNoRows {
		http.Error(w, "Урок не найден", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Error querying lesson: %v", err)
		http.Error(w, "Ошибка получения урока: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// teacherID.Valid означает, что поле не NULL
	if teacherID.Valid && teacherUsername.Valid {
		tID := int(teacherID.Int64)
		lesson.TeacherID = &tID
		lesson.Teacher = &models.Teacher{
			ID:       tID,
			Username: teacherUsername.String,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lesson)
}

// Пример GetLessonWithStudents (если нужна более подробная логика)
func GetLessonWithStudents(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID урока обязателен", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID урока", http.StatusBadRequest)
		return
	}

	var lesson models.LessonWithStudents
	var teacherID sql.NullInt64
	var teacherUsername sql.NullString

	err = db.DB.QueryRow(`
        SELECT l.id, l.name, l.description, l.teacher_id,
               t.username
        FROM lessons l
        LEFT JOIN users t ON l.teacher_id = t.id
        WHERE l.id = $1
    `, id).Scan(
		&lesson.ID, &lesson.Name, &lesson.Description,
		&teacherID, &teacherUsername,
	)

	if err == sql.ErrNoRows {
		http.Error(w, "Урок не найден", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Ошибка получения урока: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if teacherID.Valid && teacherUsername.Valid {
		tID := int(teacherID.Int64)
		lesson.Teacher = &models.Teacher{
			ID:       tID,
			Username: teacherUsername.String,
		}
		lesson.TeacherID = &tID
	}

	rows, err := db.DB.Query(`
        SELECT 
            u.id, 
            u.username, 
            ls.grade
        FROM users u
        JOIN lesson_students ls ON u.id = ls.student_id
        WHERE ls.lesson_id = $1 AND u.role = 'student'
    `, id)
	if err != nil {
		http.Error(w, "Ошибка получения списка студентов: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var student models.StudentAssignment
		var grade sql.NullInt64
		if err := rows.Scan(&student.ID, &student.Username, &grade); err != nil {
			http.Error(w, "Ошибка обработки данных студента: "+err.Error(), http.StatusInternalServerError)
			return
		}
		if grade.Valid {
			gradeInt := int(grade.Int64)
			student.Grade = &gradeInt
		}
		lesson.Students = append(lesson.Students, models.Student{
			ID:       student.ID,
			Username: student.Username,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lesson)
}
