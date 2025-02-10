package routes

import (
	"github.com/gorilla/mux"
	"mini_moodle/backend/handlers"
	"mini_moodle/backend/middleware"
	"net/http"
)

func SetupRouter(staticPath string) *mux.Router {
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/api/register", handlers.RegisterUser).Methods("POST")
	router.HandleFunc("/api/login", handlers.LoginUser).Methods("POST")

	// Protected routes
	authRouter := router.PathPrefix("/api").Subrouter()
	authRouter.Use(middleware.AuthMiddleware)

	authRouter.HandleFunc("/lesson", handlers.GetLesson).Methods("GET")
	authRouter.HandleFunc("/lessons", handlers.GetLessons).Methods("GET")
	authRouter.HandleFunc("/lesson/delete", handlers.DeleteLesson).Methods("DELETE")
	authRouter.HandleFunc("/teachers", handlers.GetTeachers).Methods("GET")
	authRouter.HandleFunc("/lesson/assign-teacher", handlers.AssignTeacher).Methods("POST")
	authRouter.HandleFunc("/students", handlers.GetStudents).Methods("GET")
	authRouter.HandleFunc("/lesson/assign-students", handlers.AssignStudents).Methods("POST")

	// Serve static files (frontend) из переданного пути
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(staticPath)))

	router.Use(mux.CORSMethodMiddleware(router))
	return router
}
