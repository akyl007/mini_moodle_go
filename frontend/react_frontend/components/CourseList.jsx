import React, { useEffect, useState } from 'react';
import axios from 'axios';

const CourseList = () => {
    const [courses, setCourses] = useState([]);

    useEffect(() => {
        fetchCourses();
    }, []);

    const fetchCourses = async () => {
        try {
            const response = await axios.get('/api/lessons');
            setCourses(response.data);
        } catch (error) {
            console.error('Ошибка загрузки курсов:', error);
        }
    };

    return (
        <div>
            <h1>Список курсов</h1>
            <ul>
                {courses.map(course => (
                    <li key={course.id}>
                        <h2>{course.name}</h2>
                        <p>{course.description}</p>
                        <p><strong>Преподаватель:</strong> {course.teacher_id || 'Не назначен'}</p>
                        <button onClick={() => navigateToLesson(course.id)}>Перейти к уроку</button>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default CourseList;