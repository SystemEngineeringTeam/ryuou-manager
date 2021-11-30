import React from "react";
import ApplicationBar from "./applicationBar";
import QuestionList from "./questionList";
import { Routes, Route, BrowserRouter } from "react-router-dom";
import Login from "./login";

const App = () => {
  return (
    <BrowserRouter>
      <ApplicationBar />
      <Routes>
        <Route path="list" element={<QuestionList />} />
        <Route path="login" element={<Login />} />
      </Routes>
    </BrowserRouter>
  );
};

export default App;
