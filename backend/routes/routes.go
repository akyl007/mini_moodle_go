package routes

import (
	"github.com/gorilla/mux"
	"mini_moodle/backend/handlers"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/lesson", handlers.GetLesson).Methods("GET")
	router.HandleFunc("/api/lessons", handlers.GetLessons).Methods("GET")
	router.HandleFunc("/api/lesson/delete", handlers.DeleteLesson).Methods("DELETE")
	router.HandleFunc("/api/teachers", handlers.GetTeachers).Methods("GET")
	router.HandleFunc("/api/lesson/assign-teacher", handlers.AssignTeacher).Methods("POST")
	router.HandleFunc("/api/students", handlers.GetStudents).Methods("GET")
	router.HandleFunc("/api/lesson/assign-students", handlers.AssignStudents).Methods("POST")

	router.Use(mux.CORSMethodMiddleware(router))

	return router
}
