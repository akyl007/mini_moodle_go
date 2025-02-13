// lesson.js
import { api } from './api.js';

document.addEventListener('DOMContentLoaded', async () => {
    const lessonTitle = document.getElementById('lesson-title');
    const lessonDescription = document.getElementById('lesson-description');
    const lessonTeacher = document.getElementById('lesson-teacher');

    const params = new URLSearchParams(window.location.search);
    const lessonId = params.get('id');

    if (!lessonId) {
        lessonTitle.textContent = 'Урок не выбран.';
        return;
    }

    try {
        const lesson = await api.getLessonDetails(lessonId);
        lessonTitle.textContent = lesson.name;
        lessonDescription.textContent = lesson.description;
        lessonTeacher.innerHTML = lesson.teacher_id
            ? `<strong>Преподаватель:</strong> ${lesson.teacher_id}`
            : `<strong>Преподаватель не назначен</strong>`;

        // Пример: добавление кнопки для открытия модального окна для назначения преподавателя
        const assignTeacherBtn = document.createElement('button');
        assignTeacherBtn.textContent = 'Назначить преподавателя';
        assignTeacherBtn.addEventListener('click', () => {
            openAssignTeacherModal(lesson.id);
        });
        // Добавляем кнопку после заголовка
        lessonTitle.parentNode.appendChild(assignTeacherBtn);
    } catch (err) {
        lessonTitle.textContent = err.message;
    }
});

function openAssignTeacherModal(lessonId) {
    // Ожидается, что в HTML есть модальное окно с id "assignTeacherModal"
    const modal = document.getElementById('assignTeacherModal');
    modal.style.display = 'block';

    // Обработчик для кнопки назначения преподавателя
    document.getElementById('assignTeacherBtn').addEventListener('click', async () => {
        const teacherSelect = document.getElementById('teacherSelect');
        const teacherId = teacherSelect.value;
        try {
            await api.assignTeacher(lessonId, teacherId);
            modal.style.display = 'none';
            alert('Преподаватель назначен успешно!');
            location.reload();
        } catch (err) {
            alert('Ошибка: ' + err.message);
        }
    });

    // Обработчик для закрытия модального окна
    document.getElementById('assignTeacherClose').addEventListener('click', () => {
        modal.style.display = 'none';
    });
}
