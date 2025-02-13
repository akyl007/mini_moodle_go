// profile.js
import { api } from './api.js';

document.addEventListener('DOMContentLoaded', async () => {
    // Логика сайдбара
    const menuToggle = document.getElementById('menu-toggle');
    const sidebar = document.getElementById('sidebar');
    menuToggle.addEventListener('click', () => {
        sidebar.classList.toggle('active');
    });

    // Логика выхода
    const logoutLink = document.getElementById('logout-link');
    logoutLink.addEventListener('click', (event) => {
        event.preventDefault();
        localStorage.removeItem('token');
        window.location.href = 'login.html';
    });

    // Логика загрузки профиля
    const profileInfo = document.getElementById('profile-info');

    try {
        // Допустим, у тебя есть эндпоинт /api/progress/student
        // или /api/user/me, где можно получить информацию о текущем пользователе
        // Пример:
        const data = await api.request('/api/progress/student', { method: 'GET' });

        // data может содержать список прогресса по урокам, курсам и т.д.
        // Можно вывести какую-то информацию о пользователе
        // Если нужно просто имя пользователя, возможно, у тебя есть другой эндпоинт
        // Здесь показываю пример, как можно вывести список прогресса
        if (Array.isArray(data) && data.length > 0) {
            let html = '<h2>Ваш прогресс:</h2>';
            data.forEach(item => {
                html += `
          <div class="progress-item">
            <strong>Урок:</strong> ${item.LessonName || 'Неизвестно'} <br>
            <strong>Оценка:</strong> ${item.Grade || 'N/A'} <br>
            <strong>Завершён:</strong> ${item.Completed ? 'Да' : 'Нет'} <br>
            <hr>
        `;
            });
            profileInfo.innerHTML = html;
        } else {
            profileInfo.innerHTML = '<p>Нет данных о прогрессе</p>';
        }
    } catch (err) {
        profileInfo.innerHTML = `<p class="error-message">Ошибка загрузки профиля: ${err.message}</p>`;
    }
});
