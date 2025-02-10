import React, {useEffect, useState} from 'react';
import axios from 'axios';

const AssignTeacherModal = ({ lessonId, onClose }) => {
    const [teachers, setTeachers] = useState([]);
    const [selectedTeacher, setSelectedTeacher] = useState('');

    useEffect(() => {
        fetchTeachers();
    }, []);

    const fetchTeachers = async () => {
        try {
            const response = await axios.get('/api/teachers');
            setTeachers(response.data);
        } catch (error) {
            console.error('Ошибка загрузки преподавателей:', error);
        }
    };

    const handleAssign = async () => {
        try {
            await axios.post('/api/lesson/assign-teacher', {
                lesson_id: lessonId,
                teacher_id: selectedTeacher
            });
            onClose();
        } catch (error) {
            console.error('Ошибка назначения преподавателя:', error);
        }
    };

    return (
        <div className="modal">
            <h2>Назначить преподавателя</h2>
            <select value={selectedTeacher} onChange={(e) => setSelectedTeacher(e.target.value)}>
                <option value="">Выберите преподавателя</option>
                {teachers.map(teacher => (
                    <option key={teacher.id} value={teacher.id}>{teacher.username}</option>
                ))}
            </select>
            <button onClick={handleAssign}>Назначить</button>
            <button onClick={onClose}>Закрыть</button>
        </div>
    );
};

export default AssignTeacherModal;