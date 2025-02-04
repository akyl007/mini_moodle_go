document.addEventListener("DOMContentLoaded", function() {
    loadLesson();
});

function loadLesson() {
    const urlParams = new URLSearchParams(window.location.search);
    const lessonId = urlParams.get("id");

    if (!lessonId) {
        document.getElementById("lesson-title").innerText = "Урок не найден";
        return;
    }

    fetch(`http://localhost:8080/api/lesson?id=${lessonId}`)
        .then(response => {
            if (!response.ok) {
                return response.text().then(text => {
                    throw new Error(`Ошибка загрузки данных: ${response.status} - ${text}`);
                });
            }
            return response.json();
        })
        .then(lesson => {
            if (!lesson || Object.keys(lesson).length === 0) {
                document.getElementById("lesson-title").innerText = "Урок не найден";
                return;
            }

            document.getElementById("lesson-title").innerText = lesson.name;
            document.getElementById("lesson-description").innerText = lesson.description;
            document.getElementById("lesson-teacher").innerHTML = `<strong>Преподаватель:</strong> ${lesson.teacher_id === null ? "Не назначен" : "ID " + lesson.teacher_id}`;
        })
        .catch(error => {
            console.error("Ошибка загрузки урока:", error);
            document.getElementById("lesson-title").innerText = "Ошибка загрузки данных";
        });
}
