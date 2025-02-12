package models

type LessonAssignment struct {
	LessonID  int  `json:"lesson_id"`
	StudentID int  `json:"student_id"`
	Grade     *int `json:"grade,omitempty"`
}

type LessonWithStudents struct {
	Lesson
	Students []Student `json:"students"`
	Teacher  *Teacher  `json:"teacher,omitempty"`
}

type StudentAssignment struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Grade    *int   `json:"grade,omitempty"`
}
