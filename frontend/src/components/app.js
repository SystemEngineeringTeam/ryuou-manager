import React from "react";
import ApplicationBar from "./applicationBar";
import QuestionList from "./questions/list/questionList";
import { Routes, Route, BrowserRouter } from "react-router-dom";
import Signup from "./users/signup";
import Login from "./users/login";
import QuestionInfo from "./questions/list/questionInfo";
import TeamCreate from "./team/teamCreate"
import TeamDivision from "./team/teamDivision"


const App = () => {
  return (
    <BrowserRouter>
      <ApplicationBar />
      <Routes>
        <Route path="list" element={<QuestionList />} />
        <Route path="signup" element={<Signup />} />
        <Route path="login" element={<Login />} />
        <Route path="questionInfo/:id" element={<QuestionInfo />} />
        <Route path="teamCreate" element={<TeamCreate />} />
        <Route path="teamDivision" element={<TeamDivision />} />
      </Routes>
    </BrowserRouter>
  );
};

export default App;
