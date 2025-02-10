import React, { useEffect, useState } from 'react';
import axios from 'axios';

const AssignStudentsModal = ({ lessonId, onClose }) => {
    const [students, setStudents] = useState([]);
    const [selectedStudents, setSelectedStudents] = useState([]);

    useEffect(() => {
        fetchStudents();
    }, []);

    const fetchStudents = async () => {
        try {
            const response = await axios.get('/api/students');
            setStudents(response.data);
        } catch (error) {
            console.error('Ошибка загрузки студентов:', error);
        }
    };

    const handleAssign = async () => {
        try {
            await axios.post('/api/lesson/assign-students', {
                lesson_id: lessonId,
                student_ids: selectedStudents
            });
            onClose();
        } catch (error) {
            console.error('Ошибка назначения студентов:', error);
        }
    };

    return (
        <div className="modal">
            <h2>Назначить студентов</h2>
            <select multiple value={selectedStudents} onChange={(e) => setSelectedStudents(Array.from(e.target.selectedOptions, option => option.value))}>
                {students.map(student => (
                    <option key={student.id} value={student.id}>{student.username}</option>
                ))}
            </select>
            <button onClick={handleAssign}>Назначить</button>
            <button onClick={onClose}>Закрыть</button>
        </div>
    );
};

export default AssignStudentsModal;