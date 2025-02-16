import { api } from './api.js';
import { checkAuth } from './auth.js';

document.addEventListener('DOMContentLoaded', async () => {
    if (!checkAuth()) return;

    const menuToggle = document.getElementById('menu-toggle');
    const sidebar = document.getElementById('sidebar');
    menuToggle.addEventListener('click', () => sidebar.classList.toggle('active'));

    const logoutLink = document.getElementById('logout-link');
    logoutLink.addEventListener('click', (event) => {
        event.preventDefault();
        localStorage.removeItem('token');
        localStorage.removeItem('role');
        window.location.href = 'login.html';
    });

    const params = new URLSearchParams(window.location.search);
    const lessonId = params.get('lesson_id');
    const lessonTitle = document.getElementById('lesson-title');
    const lessonDesc = document.getElementById('lesson-description');
    const lessonTeacher = document.getElementById('lesson-teacher');

    try {
        const lesson = await api.getLessonDetails(lessonId);
        lessonTitle.textContent = lesson.name;
        lessonDesc.textContent = lesson.description;
        if (lesson.teacher && lesson.teacher.username) {
            lessonTeacher.innerHTML = `<strong>Преподаватель:</strong> ${lesson.teacher.username}`;
        } else {
            lessonTeacher.innerHTML = `<strong>Преподаватель не назначен</strong>`;
        }
    } catch (err) {
        lessonTitle.textContent = `Ошибка: ${err.message}`;
    }
});
