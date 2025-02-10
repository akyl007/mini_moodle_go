import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { useParams } from 'react-router-dom';

const LessonDetail = () => {
    const { id } = useParams();
    const [lesson, setLesson] = useState(null);

    useEffect(() => {
        fetchLesson();
    }, [id]);

    const fetchLesson = async () => {
        try {
            const response = await axios.get(`/api/lesson?id=${id}`);
            setLesson(response.data);
        } catch (error) {
            console.error('Ошибка загрузки урока:', error);
        }
    };

    if (!lesson) {
        return <div>Загрузка...</div>;
    }

    return (
        <div>
            <h1>{lesson.name}</h1>
            <p>{lesson.description}</p>
            <p><strong>Преподаватель:</strong> {lesson.teacher_id || 'Не назначен'}</p>
        </div>
    );
};

export default LessonDetail;