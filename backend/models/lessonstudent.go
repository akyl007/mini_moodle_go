package models

type LessonStudent struct {
	ID        int `json:"id"`
	LessonID  int `json:"lesson_id"`
	StudentID int `json:"student_id"`
}
