package routes

import (
	"github.com/gorilla/mux"
	"mini_moodle/backend/handlers"
	"mini_moodle/backend/middleware"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Публичные маршруты
	router.HandleFunc("/api/login", handlers.Login).Methods("POST")
	router.HandleFunc("/api/register", handlers.Register).Methods("POST")

	// Защищенные маршруты
	// Маршруты для преподавателей и администраторов
	router.HandleFunc("/api/lesson", middleware.TeacherOrAdmin(handlers.CreateLesson)).Methods("POST")
	router.HandleFunc("/api/lesson/delete", middleware.TeacherOrAdmin(handlers.DeleteLesson)).Methods("DELETE")
	router.HandleFunc("/api/lesson/assign-teacher", middleware.TeacherOrAdmin(handlers.AssignTeacher)).Methods("POST")
	router.HandleFunc("/api/lesson/assign-students", middleware.TeacherOrAdmin(handlers.AssignStudents)).Methods("POST")
	router.HandleFunc("/api/lesson/grade", middleware.TeacherOrAdmin(handlers.AssignGrade)).Methods("POST")
	router.HandleFunc("/api/course", middleware.TeacherOrAdmin(handlers.CreateCourse)).Methods("POST")
	router.HandleFunc("/api/course", middleware.TeacherOrAdmin(handlers.UpdateCourse)).Methods("PUT")
	router.HandleFunc("/api/course", middleware.TeacherOrAdmin(handlers.DeleteCourse)).Methods("DELETE")

	// Обновленные маршруты для назначения
	router.HandleFunc("/api/course/assign-teacher", middleware.TeacherOrAdmin(handlers.AssignTeacher)).Methods("POST")
	router.HandleFunc("/api/course/assign-students", middleware.TeacherOrAdmin(handlers.AssignStudents)).Methods("POST")
	
	// Маршрут для посещаемости
	router.HandleFunc("/api/lesson/attendance", middleware.TeacherOrAdmin(handlers.UpdateAttendance)).Methods("POST")

	// Маршруты для всех аутентифицированных пользователей
	router.HandleFunc("/api/teachers", middleware.AuthMiddleware(handlers.GetTeachers)).Methods("GET")
	router.HandleFunc("/api/students", middleware.AuthMiddleware(handlers.GetStudents)).Methods("GET")
	router.HandleFunc("/api/courses", middleware.AuthMiddleware(handlers.GetCourses)).Methods("GET")
	router.HandleFunc("/api/lessons", middleware.AuthMiddleware(handlers.GetLessons)).Methods("GET")
	router.HandleFunc("/api/lesson", middleware.AuthMiddleware(handlers.GetLesson)).Methods("GET")
	router.HandleFunc("/api/progress/student", middleware.AuthMiddleware(handlers.GetStudentProgress)).Methods("GET")
	router.HandleFunc("/api/progress/course", middleware.TeacherOrAdmin(handlers.GetCourseProgress)).Methods("GET")
	router.HandleFunc("/api/courses/{id}", handlers.GetCourse).Methods("GET")

	// Feedback route
	router.HandleFunc("/api/feedback", middleware.AuthMiddleware(handlers.CreateFeedback)).Methods("POST")

	router.Use(mux.CORSMethodMiddleware(router))

	return router
}
