import { api } from './api.js';

function validateRegistration(username, password) {
    if (username.length < 3) {
        return "Имя пользователя должно быть не менее 3 символов.";
    }
    const passwordRegex = /^(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*])[A-Za-z\d!@#$%^&*]{8,}$/;
    if (!passwordRegex.test(password)) {
        return "Пароль должен быть не менее 8 символов, содержать одну цифру, одну заглавную букву и один спецсимвол.";
    }
    return "";
}

document.addEventListener('DOMContentLoaded', () => {
    const registerForm = document.getElementById('register-form');
    const errorMessageEl = document.getElementById('error-message');

    registerForm.addEventListener('submit', async (event) => {
        event.preventDefault();
        errorMessageEl.style.display = 'none';
        errorMessageEl.textContent = '';

        const username = document.getElementById('username').value.trim();
        const password = document.getElementById('password').value;
        const role = "student"; // Фиксированная роль для регистрации студента

        const validationError = validateRegistration(username, password);
        if (validationError) {
            errorMessageEl.textContent = validationError;
            errorMessageEl.style.display = 'block';
            return;
        }

        try {
            const data = await api.register(username, password, role);
            window.location.href = 'login.html';
        } catch (err) {
            errorMessageEl.textContent = err.message;
            errorMessageEl.style.display = 'block';
        }
    });
});
