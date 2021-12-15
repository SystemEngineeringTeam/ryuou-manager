import React from "react";
import ApplicationBar from "./applicationBar";
import QuestionList from "./questions/list/questionList";
import { Routes, Route, BrowserRouter } from "react-router-dom";
import Signup from "./users/signup";
import Login from "./users/login";
import QuestionInfo from "./questions/list/questionInfo";
import CreateForm from "./admin/questions/CreateForm";
import SubmitList from "./admin/questions/SubmitList";

const App = () => {
  return (
    <BrowserRouter>
      <ApplicationBar />
      <Routes>
        <Route path="list" element={<QuestionList />} />
        <Route path="signup" element={<Signup />} />
        <Route path="login" element={<Login />} />
        <Route path="questionInfo/:id" element={<QuestionInfo />} />
        <Route path="admin/questions/create" element={<CreateForm />} />
        <Route path="admin/submits" element={<SubmitList />} />
      </Routes>
    </BrowserRouter>
  );
};

export default App;
