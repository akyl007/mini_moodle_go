import { api } from './api.js';
import { checkAuth, getUserRole } from './auth.js';

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

    const profileInfo = document.getElementById('profile-info');

    // Для простоты, выводим имя и роль пользователя из localStorage
    const username = localStorage.getItem('username') || "Неизвестно";
    const role = getUserRole() || "Неизвестно";

    // В будущем можно добавить дополнительные данные, например, прогресс по курсу и т.д.
    profileInfo.innerHTML = `
    <p><strong>Имя пользователя:</strong> ${username}</p>
    <p><strong>Роль:</strong> ${role}</p>
  `;

    // Дополнительно: если пользователь студент, можно вызвать API для получения прогресса
    if (role === "student") {
        try {
            const progress = await api.request('/api/progress/student', { method: 'GET' });
            // Обработка и вывод прогресса – пример
            if (Array.isArray(progress) && progress.length > 0) {
                let progressHTML = '<h2>Ваш прогресс:</h2>';
                progress.forEach(item => {
                    progressHTML += `
            <div class="progress-item">
              <strong>Урок:</strong> ${item.lesson_name || 'Неизвестно'}<br>
              <strong>Оценка:</strong> ${item.grade || 'N/A'}<br>
              <strong>Посещаемость:</strong> ${item.completed ? 'Да' : 'Нет'}<br>
              <hr>
            </div>
          `;
                });
                profileInfo.innerHTML += progressHTML;
            } else {
                profileInfo.innerHTML += `<p>Данных о прогрессе нет.</p>`;
            }
        } catch (err) {
            profileInfo.innerHTML += `<p class="error-message">Ошибка загрузки прогресса: ${err.message}</p>`;
        }
    }
});
