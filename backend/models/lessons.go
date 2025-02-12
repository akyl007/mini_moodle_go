package models

type Lesson struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CourseID    int    `json:"course_id"`
	TeacherID   *int   `json:"teacher_id,omitempty"`
}

type LessonWithTeacher struct {
	Lesson
	Teacher *Teacher `json:"teacher,omitempty"`
}
