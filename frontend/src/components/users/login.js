import { Button, TextField, Grid, Paper } from "@mui/material";
import { Box } from "@mui/system";
import axios from "axios";
import React from "react";
import Path from "../.react.config";
import { useCookies } from "react-cookie";
import { useNavigate } from "react-router-dom";

const Login = () => {
  const [cookie, setCookie] = useCookies();
  const email = React.useRef();
  const password = React.useRef();

  const navigate = useNavigate();

  const doSubmit = async (e) => {
    e.preventDefault();
    const res = await axios.post(Path.Login, {
      email: email.current.value,
      password: password.current.value,
    });
    if (res.status !== 200) {
      alert("ログインに失敗しました");
      return;
    }

    console.log(res.data);
    setCookie("userID", res.data.user_id);
    setCookie("teamID", res.data.team_id);
    setCookie("sessionID", res.data.session_id);

    if (res.data.session_id == "admin") {
      navigate("/admin");
    } else {
      navigate("/list");
    }
  };

  console.log(cookie.userID);

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
                inputRef={password}
                margin="normal"
                label="password"
              />
            </Box>

            <Box sx={{ mx: "6rem", my: "2rem" }}>
              <Button type="submit" variant="contained">
                ログイン
              </Button>
            </Box>
          </form>
        </Grid>
      </Paper>
    </Box>
  );
};

export default Login;
