import { Button, TextField } from "@mui/material";
import { Box } from "@mui/system";
import axios from "axios";
import React from "react";
import Path from "./.react.config";

const Login = () => {
  const doSubmit = async (e) => {
    e.preventDefault();
    const res = await axios.post(Path.Login);
    console.log(res);
  };

  return (
    <Box>
      <h1>Login</h1>
      <form onSubmit={doSubmit}>
        <TextField placeholder="UserName" type="email" />
        <TextField placeholder="password" type="password" />
        <Button type="submit" variant="contained">
          Login
        </Button>
      </form>
    </Box>
  );
};

export default Login;
