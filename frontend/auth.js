export function checkAuth(noticeElId = 'notice') {
    const token = localStorage.getItem('token');
    const noticeEl = document.getElementById(noticeElId);
    if (!token) {
        if (noticeEl) {
            noticeEl.innerHTML = "Для просмотра содержимого необходимо авторизоваться. <a href='login.html'>Войти</a>";
            noticeEl.style.display = 'block';
        }
        return false;
    }
    if (noticeEl) {
        noticeEl.style.display = 'none';
    }
    return true;
}

export function getUserRole() {
    return localStorage.getItem('role');
}
