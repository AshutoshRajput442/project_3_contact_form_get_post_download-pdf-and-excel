import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Form from "./components/Form";
import SubmittedData from "./components/SubmittedData";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Form />} />
        <Route path="/submitted" element={<SubmittedData />} />
      </Routes>
    </Router>
  );
}

export default App;
