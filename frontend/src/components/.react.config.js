const baseURL = "http://localhost:8080";

const Path = {
  Question: baseURL + "/questions",
  User: baseURL + "/users",
  Login: baseURL + "/users/login",
  AdminQuestion: baseURL + "/admin/questions",
  Admin: {
    User: baseURL + "/admin/users",
    Team: baseURL + "/admin/teams",
  },
};
export default Path;
