-- Создаем enum для ролей
CREATE TYPE user_role AS ENUM ('admin', 'teacher', 'student');

-- Обновляем таблицу пользователей
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role user_role NOT NULL DEFAULT 'student'
);

CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE lessons (
    id SERIAL PRIMARY KEY,
    course_id INTEGER REFERENCES courses(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    teacher_id INTEGER REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE lesson_students (
    lesson_id INTEGER REFERENCES lessons(id) ON DELETE CASCADE,
    student_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    grade INTEGER,
    PRIMARY KEY (lesson_id, student_id)
);

-- Добавляем индексы для оптимизации запросов прогресса
CREATE INDEX idx_lesson_students_student_id ON lesson_students(student_id);
CREATE INDEX idx_lesson_students_lesson_id ON lesson_students(lesson_id);
CREATE INDEX idx_lessons_course_id ON lessons(course_id); 