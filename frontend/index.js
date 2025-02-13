// index.js
import { api } from './api.js';

document.addEventListener('DOMContentLoaded', async () => {
    // Обработка кнопки для открытия/закрытия сайдбара
    const menuToggle = document.getElementById('menu-toggle');
    const sidebar = document.getElementById('sidebar');
    menuToggle.addEventListener('click', () => {
        sidebar.classList.toggle('active');
    });

    // Загрузка списка уроков
    const lessonsContainer = document.getElementById('lessons-container');
    try {
        const lessons = await api.getLessons();
        lessonsContainer.innerHTML = ''; // Очищаем контейнер

        if (lessons.length === 0) {
            lessonsContainer.innerHTML = '<p>Нет уроков для отображения.</p>';
        } else {
            lessons.forEach(lesson => {
                const card = document.createElement('div');
                card.className = 'lesson-card';
                card.innerHTML = `
          <h3>${lesson.name}</h3>
          <p>${lesson.description}</p>
        `;
                // Переход на страницу деталей урока по клику
                card.addEventListener('click', () => {
                    window.location.href = `lesson.html?id=${lesson.id}`;
                });
                lessonsContainer.appendChild(card);
            });
        }
    } catch (err) {
        lessonsContainer.innerHTML = `<p class="error-message">${err.message}</p>`;
    }
});
