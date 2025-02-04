package models

type Lesson struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TeacherID   *int   `json:"teacher_id,omitempty"`
}

type Teacher struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type Student struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}
