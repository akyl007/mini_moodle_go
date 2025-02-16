document.addEventListener('DOMContentLoaded', () => {
    const logoutLink = document.getElementById('logout-link');
    if (logoutLink) {
        logoutLink.addEventListener('click', (event) => {
            event.preventDefault();
            localStorage.removeItem('token');
            localStorage.removeItem('role');
            window.location.href = 'login.html';
        });
    }
});