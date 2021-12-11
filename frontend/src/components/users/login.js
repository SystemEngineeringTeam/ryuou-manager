import { Button, TextField } from "@mui/material";
import { Box } from "@mui/system";
import axios from "axios";
import React from "react";
import Path from "../.react.config";

const Login = () => {
  const email = React.useRef();
  const password = React.useRef();

  const doSubmit = async (e) => {
    e.preventDefault();
    const res = await axios.post(Path.Login, {
      email: email.current.value,
      password: password.current.value,
    });
    console.log(res);
  };

  return (
    <Box>
      <h1>Login</h1>
      <form onSubmit={doSubmit}>
        <TextField placeholder="email" type="email" inputRef={email} />
        <TextField placeholder="password" type="password" inputRef={password} />
        <Button type="submit" variant="contained">
          Login
        </Button>
      </form>
    </Box>
  );
};

export default Login;
