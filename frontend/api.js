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

    // Получение данных
    getCourses() {
        return this.request('/api/courses', { method: 'GET' });
    }

    // Получить уроки для конкретного курса
    getLessonsByCourse(courseId) {
        return this.request(`/api/lessons?course_id=${courseId}`, { method: 'GET' });
    }

    // Получить детали конкретного урока
    getLessonDetails(lessonId) {
        return this.request(`/api/lesson?id=${lessonId}`, { method: 'GET' });
    }
    getCourseDetails(courseId) {
        return this.request(`/api/courses/${courseId}`, { method: 'GET' }); // Используем endpoint /api/courses/{id}
    }
    getTeachers() {
        return this.request('/api/teachers', { method: 'GET' });
    }

    getStudents() {
        return this.request('/api/students', { method: 'GET' });
    }

    // Операции назначения (для курсов – преподаватель, для курса – студентов)
    // Предполагаем, что назначение преподавателя происходит на уровне курса:
    assignTeacher(courseId, teacherId) {
        return this.request('/api/course/assign-teacher', {
            method: 'POST',
            body: JSON.stringify({
                course_id: parseInt(courseId),
                teacher_id: parseInt(teacherId)
            }),
        });
    }

    assignStudents(courseId, studentIds) {
        return this.request('/api/course/assign-students', {
            method: 'POST',
            body: JSON.stringify({ course_id: courseId, student_ids: studentIds }),
        });
    }
}

export const api = new API('http://localhost:8080');
