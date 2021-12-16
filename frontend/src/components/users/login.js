import { Button, TextField, Grid, Paper } from "@mui/material";
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
    <Box
      sx={{
        m: 1,
        width: 800,
        height: 450,
        margin: "auto",
      }}
    >
      <Paper variant="outlined">
        <Grid container justifyContent="center" alignItems="center">
          <form onSubmit={doSubmit}>
            <Box
              sx={{
                display: "flex",
                flexDirection: "column",
                "& .MuiTextField-root": { width: "25ch" },
              }}
            >
              <h1>Login</h1>
              <TextField
                placeholder="email"
                type="email"
                inputRef={email}
                margin="normal"
                label="email"
              />
              <TextField
                placeholder="password"
                type="password"
                ninputRef={password}
                margin="normal"
                label="password"
              />
            </Box>

            <Box sx={{ mx: "6rem", my: "2rem" }}>
              <Button type="submit" variant="contained">
                Login
              </Button>
            </Box>
          </form>
        </Grid>
      </Paper>
    </Box>
  );
};

export default Login;
