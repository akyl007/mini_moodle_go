document.addEventListener("DOMContentLoaded", function() {
    fetchLessons();
});

let selectedLessonId = null;

function toggleSidebar() {
    const sidebar = document.getElementById("sidebar");
    const overlay = document.getElementById("overlay");

    // Переключаем классы для панели и затемнения
    sidebar.classList.toggle("active");
    overlay.classList.toggle("active");

    // Если меню закрывается, скрываем затемнение с задержкой
    if (!sidebar.classList.contains("active")) {
        setTimeout(() => {
            overlay.style.display = "none";
        }, 300);
    } else {
        overlay.style.display = "block";
    }
}
function fetchLessons() {
    fetch("http://localhost:8080/api/lessons")
        .then(response => response.json())
        .then(data => {
            console.log("Полученные данные:", data);
            const container = document.getElementById("lessons-container");
            container.innerHTML = "";

            data.forEach(lesson => {
                const div = document.createElement("div");
                div.className = "lesson";
                div.innerHTML = `
                    <h3>${lesson.name}</h3>
                    <p>${lesson.description}</p>
                    <p><strong>Преподаватель:</strong> ${lesson.teacher_id === null ? "Не назначен" : "ID " + lesson.teacher_id}</p>
                    <button class="assign-teacher-btn" onclick="openAssignTeacherModal(event, ${lesson.id})">📘 Назначить преподавателя</button>
                    <button class="assign-student-btn" onclick="openAssignStudentModal(event, ${lesson.id})">🎓 Назначить студентов</button>
                `;

                div.addEventListener("click", function() {
                    window.location.href = `lesson.html?id=${lesson.id}`;
                });

                container.appendChild(div);
            });
        })
        .catch(error => console.error("Ошибка:", error));
}


function deleteLesson(id) {
    fetch(`http://localhost:8080/api/lesson/delete?id=${id}`, { method: "DELETE" })
        .then(response => response.text())
        .then(() => fetchLessons())
        .catch(error => console.error("Ошибка удаления:", error));
}

function openAssignTeacherModal(event, lessonId) {
    event.stopPropagation(); // Останавливаем всплытие клика
    selectedLessonId = lessonId;
    const modal = document.getElementById("assignTeacherModal");
    modal.style.display = "block";

    fetch("http://localhost:8080/api/teachers")
        .then(response => response.json())
        .then(teachers => {
            const teacherSelect = document.getElementById("teacherSelect");
            teacherSelect.innerHTML = "";

            teachers.forEach(teacher => {
                const option = document.createElement("option");
                option.value = teacher.id;
                option.textContent = teacher.username;
                teacherSelect.appendChild(option);
            });
        })
        .catch(error => console.error("Ошибка загрузки преподавателей:", error));
}


function closeAssignTeacherModal() {
    document.getElementById("assignTeacherModal").style.display = "none";
}

function assignTeacher() {
    const teacherId = document.getElementById("teacherSelect").value;

    fetch("http://localhost:8080/api/lesson/assign-teacher", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ lesson_id: selectedLessonId, teacher_id: teacherId })
    })
        .then(response => response.json())
        .then(() => {
            closeAssignTeacherModal();
            fetchLessons();
        })
        .catch(error => console.error("Ошибка назначения преподавателя:", error));
}

function openAssignStudentModal(event, lessonId) {
    event.stopPropagation(); // Останавливаем всплытие клика
    selectedLessonId = lessonId;
    const modal = document.getElementById("assignStudentModal");
    modal.style.display = "block";

    fetch("http://localhost:8080/api/students")
        .then(response => response.json())
        .then(students => {
            const studentSelect = document.getElementById("studentSelect");
            studentSelect.innerHTML = "";

            students.forEach(student => {
                const option = document.createElement("option");
                option.value = student.id;
                option.textContent = student.username;
                studentSelect.appendChild(option);
            });
        })
        .catch(error => console.error("Ошибка загрузки студентов:", error));
}

function closeAssignStudentModal() {
    document.getElementById("assignStudentModal").style.display = "none";
}

function assignStudents() {
    const selectedStudents = Array.from(document.getElementById("studentSelect").selectedOptions)
        .map(option => parseInt(option.value));

    fetch("http://localhost:8080/api/lesson/assign-students", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ lesson_id: selectedLessonId, student_ids: selectedStudents })
    })
        .then(response => response.json())
        .then(() => {
            closeAssignStudentModal();
            fetchLessons(); // Обновляем список уроков
        })
        .catch(error => console.error("Ошибка назначения студентов:", error));
}

