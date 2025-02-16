import { api } from './api.js';

document.addEventListener('DOMContentLoaded', () => {
    const loginForm = document.getElementById('login-form');
    const errorMessage = document.getElementById('error-message');

    loginForm.addEventListener('submit', async (event) => {
        event.preventDefault();
        errorMessage.style.display = 'none';
        errorMessage.textContent = '';

        const username = document.getElementById('username').value.trim();
        const password = document.getElementById('password').value;

        try {
            const data = await api.login(username, password);
            localStorage.setItem('token', data.token);
            localStorage.setItem('role', data.role);
            window.location.href = 'index.html';
        } catch (err) {
            errorMessage.textContent = err.message;
            errorMessage.style.display = 'block';
        }
    });
});
