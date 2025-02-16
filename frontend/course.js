import { api } from './api.js';
import { checkAuth, getUserRole } from './auth.js';

document.addEventListener('DOMContentLoaded', async () => {
    if (!checkAuth()) return;

    const menuToggle = document.getElementById('menu-toggle');
    const sidebar = document.getElementById('sidebar');
    menuToggle.addEventListener('click', () => sidebar.classList.toggle('active'));

    // Считываем course_id из URL
    const params = new URLSearchParams(window.location.search);
    const courseId = params.get('course_id');

    const courseTitle = document.getElementById('course-title');
    const courseDesc = document.getElementById('course-description');
    const lessonsContainer = document.getElementById('lessons-container');

    // Предположим, что мы просто выводим ID курса (если нет эндпоинта для подробностей)
    courseTitle.textContent = `Уроки курса #${courseId}`;

    try {
        const lessons = await api.getLessonsByCourse(courseId);
        lessonsContainer.innerHTML = '';
        if (!Array.isArray(lessons) || lessons.length === 0) {
            lessonsContainer.innerHTML = '<p>Нет уроков для отображения.</p>';
        } else {
            lessons.forEach(lesson => {
                const card = document.createElement('div');
                card.className = 'lesson-card';
                card.innerHTML = `
          <h3>${lesson.name}</h3>
          <p>${lesson.description || ''}</p>
        `;
                card.addEventListener('click', () => {
                    window.location.href = `lesson.html?course_id=${courseId}&lesson_id=${lesson.id}`;
                });
                lessonsContainer.appendChild(card);
            });
        }
    } catch (err) {
        lessonsContainer.innerHTML = `<p class="error-message">${err.message}</p>`;
    }

    // Показываем кнопку назначения преподавателя, если роль admin
    const role = getUserRole();
    const assignTeacherCourseBtn = document.getElementById('assignTeacherCourseBtn');
    if (role === 'admin') {
        assignTeacherCourseBtn.style.display = 'block';
        assignTeacherCourseBtn.addEventListener('click', () => {
            openAssignTeacherModal(courseId);
        });
    }
});

async function openAssignTeacherModal(courseId) {
    const modal = document.getElementById('assignTeacherModal');
    modal.style.display = 'flex';
    await loadTeachersToSelect(); // Загрузка преподавателей (как и было)

    // 1. Получаем детальную информацию о курсе, чтобы узнать назначенных студентов
    const courseDetails = await api.getCourseDetails(courseId);
    const assignedStudentIds = courseDetails.students.map(student => student.id); // ID уже назначенных студентов
    // 2. Загружаем список всех студентов
    const allStudents = await api.getStudents();
    // 3. Контейнер для чекбоксов студентов
    const studentsCheckboxesContainer = document.getElementById('studentsCheckboxesContainer');
    studentsCheckboxesContainer.innerHTML = ''; // Очищаем контейнер перед заполнением
    allStudents.forEach(student => {
        const checkbox = document.createElement('input');
        checkbox.type = 'checkbox';
        checkbox.id = `studentCheckbox_${student.id}`;
        checkbox.value = student.id;
        // Проверяем, назначен ли студент на курс, и чекаем чекбокс, если да
        if (assignedStudentIds.includes(student.id)) {
            checkbox.checked = true;
        }
        const label = document.createElement('label');
        label.htmlFor = `studentCheckbox_${student.id}`;
        label.textContent = student.username;
        const studentDiv = document.createElement('div'); // Для форматирования каждого студента в строке
        studentDiv.appendChild(checkbox);
        studentDiv.appendChild(label);
        studentsCheckboxesContainer.appendChild(studentDiv);
    });
    // 4. Логика для чекбокса "Выбрать всех"
    const selectAllCheckbox = document.getElementById('selectAllStudentsCheckbox');
    selectAllCheckbox.addEventListener('click', () => {
        const studentCheckboxes = studentsCheckboxesContainer.querySelectorAll('input[type="checkbox"]');
        studentCheckboxes.forEach(checkbox => {
            checkbox.checked = selectAllCheckbox.checked; // Синхронизируем состояние всех чекбоксов с "Выбрать всех"
        });
    });
    // 5. Обработчик для кнопки "Назначить / Сохранить" (модифицированный)
    document.getElementById('assignTeacherBtn').onclick = async () => {
        const teacherSelect = document.getElementById('teacherSelect');
        const teacherId = teacherSelect.value;
        // Собираем ID выбранных студентов
        const selectedStudentIds = [];
        const studentCheckboxes = studentsCheckboxesContainer.querySelectorAll('input[type="checkbox"]:checked');
        studentCheckboxes.forEach(checkbox => {
            selectedStudentIds.push(parseInt(checkbox.value)); // Преобразуем value в число и добавляем в массив
        });
        try {
            await api.assignTeacher(courseId, teacherId); // Назначаем преподавателя (как и было)
            await api.assignStudents(courseId, selectedStudentIds); // Назначаем студентов (новый вызов API)
            modal.style.display = 'none';
            alert('Назначение преподавателя и студентов успешно сохранено!');
            window.location.reload();
        } catch (err) {
            alert('Ошибка: ' + err.message);
        }
    };
    // 6. Обработчик для кнопки "Отменить" (новый)
    document.getElementById('assignCancelBtn').onclick = () => {
        modal.style.display = 'none'; // Просто закрываем модальное окно
    };

    document.getElementById('assignTeacherClose').onclick = () => { // Закрытие по крестику (как и было)
        modal.style.display = 'none';
    };
}
async function loadTeachersToSelect() {
    try {
        const teachers = await api.getTeachers();
        const teacherSelect = document.getElementById('teacherSelect');
        teacherSelect.innerHTML = '';
        teachers.forEach(t => {
            const option = document.createElement('option');
            option.value = t.id;
            option.textContent = t.username; // здесь можно добавить фамилию, если есть
            teacherSelect.appendChild(option);
        });
    } catch (err) {
        alert('Ошибка при загрузке преподавателей: ' + err.message);
    }
}
