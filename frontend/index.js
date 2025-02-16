import { api } from './api.js';
import { checkAuth } from './auth.js';

document.addEventListener('DOMContentLoaded', async () => {
    if (!checkAuth()) return;

    // Обработка меню
    const menuToggle = document.getElementById('menu-toggle');
    const sidebar = document.getElementById('sidebar');
    menuToggle.addEventListener('click', () => {
        sidebar.classList.toggle('active');
    });

    // Загрузка курсов
    const container = document.getElementById('courses-container');
    try {
        const courses = await api.getCourses();
        container.innerHTML = '';
        if (!Array.isArray(courses) || courses.length === 0) {
            container.innerHTML = '<p>Нет предметов для отображения.</p>';
        } else {
            courses.forEach(course => {
                const card = document.createElement('div');
                card.className = 'course-card';
                card.innerHTML = `
          <h3>${course.name}</h3>
          <p>${course.description || ''}</p>
        `;
                card.addEventListener('click', () => {
                    window.location.href = `course.html?course_id=${course.id}`;
                });
                container.appendChild(card);
            });
        }
    } catch (err) {
        container.innerHTML = `<p class="error-message">${err.message}</p>`;
    }
});
