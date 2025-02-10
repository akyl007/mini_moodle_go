import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import CourseList from './CourseList';
import LessonDetail from './LessonDetail';

function App() {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<CourseList />} />
                <Route path="/lesson/:id" element={<LessonDetail />} />
            </Routes>
        </Router>
    );
}

export default App;