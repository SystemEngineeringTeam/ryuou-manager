import React from "react";
import ApplicationBar from "./applicationBar";
import QuestionList from "./questionList";
import { Routes, Route, BrowserRouter } from "react-router-dom";

const App = () => {
  return (
    <BrowserRouter>
      <ApplicationBar />
      <Routes>
        <Route path="list" element={<QuestionList />} />
        <Route path="login" element={<div>Login</div>} />
      </Routes>
    </BrowserRouter>
  );
};

export default App;
