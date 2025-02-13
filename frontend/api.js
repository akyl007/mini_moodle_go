// api.js
class API {
    constructor(baseUrl) {
        this.baseUrl = baseUrl;
    }

    async request(endpoint, options = {}) {
        const token = localStorage.getItem('token');
        const headers = options.headers || {};

        if (token) {
            headers['Authorization'] = `Bearer ${token}`;
        }

        options.headers = {
            'Content-Type': 'application/json',
            ...headers,
        };

        const response = await fetch(`${this.baseUrl}${endpoint}`, options);
        if (!response.ok) {
            const errorText = await response.text();
            throw new Error(`Ошибка ${response.status}: ${errorText}`);
        }
        return response.json();
    }

    // Аутентификация
    login(username, password) {
        return this.request('/api/login', {
            method: 'POST',
            body: JSON.stringify({ username, password }),
        });
    }

    register(username, password, role) {
        return this.request('/api/register', {
            method: 'POST',
            body: JSON.stringify({ username, password, role }),
        });
    }

    // Получение списков и деталей
    getLessons() {
        return this.request('/api/lessons', { method: 'GET' });
    }

    getLessonDetails(lessonId) {
        return this.request(`/api/lesson?id=${lessonId}`, { method: 'GET' });
    }

    getCourses() {
        return this.request('/api/courses', { method: 'GET' });
    }

    getTeachers() {
        return this.request('/api/teachers', { method: 'GET' });
    }

    getStudents() {
        return this.request('/api/students', { method: 'GET' });
    }

    // Операции назначения и оценок
    assignTeacher(lessonId, teacherId) {
        return this.request('/api/lesson/assign-teacher', {
            method: 'POST',
            body: JSON.stringify({ lesson_id: lessonId, teacher_id: teacherId }),
        });
    }

    assignStudents(lessonId, studentIds) {
        return this.request('/api/lesson/assign-students', {
            method: 'POST',
            body: JSON.stringify({ lesson_id: lessonId, student_ids: studentIds }),
        });
    }

    assignGrade(lessonId, studentId, grade) {
        return this.request('/api/lesson/grade', {
            method: 'POST',
            body: JSON.stringify({ lesson_id: lessonId, student_id: studentId, grade }),
        });
    }
}

export const api = new API('http://localhost:8080');
