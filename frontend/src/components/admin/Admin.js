import React from "react";
import { Button } from "@mui/material";
import { Box } from "@mui/system";
import { Link } from "react-router-dom";

const Admin = () => {
  return (
    <Box
      sx={{
        m: 1,
        width: 800,
        height: 450,
        margin: "auto",
      }}
    >
      <h1>申し訳程度の管理者ページ</h1>

      <Button variant="contained" component={Link} to="/admin/question/create">
        問題追加
      </Button>
      <Button variant="contained" component={Link} to="/admin/team/members">
        チーム一覧
      </Button>
      <Button variant="contained" component={Link} to="/admin/team/create">
        チーム追加
      </Button>
      <Button variant="contained" component={Link} to="/admin/team/division">
        チームメンバー編集
      </Button>
      <Button variant="contained" component={Link} to="/admin/submits">
        回答一覧
      </Button>
      {/* <Route path="admin/team/create" element={<TeamCreate />} />
        <Route path="admin/team/division" element={<TeamDivision />} />
        <Route path="admin/team/members" element={<MembersList />} />
        <Route path="admin/question/create" element={<CreateForm />} />
        <Route path="admin/submits" element={<SubmitList />} /> */}
    </Box>
  );
};

export default Admin;
