import React from "react";
import ApplicationBar from "./applicationBar";
import QuestionList from "./questions/list/questionList";
import { Routes, Route, BrowserRouter } from "react-router-dom";
import Signup from "./users/signup";
import Login from "./users/login";
import QuestionInfo from "./questions/list/questionInfo";
import TeamCreate from "./admin/team/teamCreate";
import TeamDivision from "./admin/team/teamDivision";
import CreateForm from "./admin/questions/CreateForm";
import SubmitList from "./admin/questions/SubmitList";
import MembersList from "./admin/team/MembersList";
import Ranking from "./ranking/Ranking";
import Redirect from "./Redirect";
import Admin from "./admin/Admin";

const App = () => {
  return (
    <BrowserRouter>
      <ApplicationBar />
      <Routes>
        <Route path="list" element={<QuestionList />} />
        <Route path="signup" element={<Signup />} />
        <Route path="login" element={<Login />} />
        <Route path="questionInfo/:id" element={<QuestionInfo />} />
        <Route path="admin" element={<Admin />} />
        <Route path="admin/team/create" element={<TeamCreate />} />
        <Route path="admin/team/division" element={<TeamDivision />} />
        <Route path="admin/team/members" element={<MembersList />} />
        <Route path="admin/question/create" element={<CreateForm />} />
        <Route path="admin/submits" element={<SubmitList />} />
        <Route path="ranking" element={<Ranking />} />
        <Route path="*" element={<Redirect />} />
      </Routes>
    </BrowserRouter>
  );
};

export default App;
