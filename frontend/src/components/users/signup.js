import React from "react";
import { Button, TextField, Grid, Paper, Box } from "@mui/material";
import axios from "axios";
import Path from "../.react.config";

const Signup = () => {
  const name = React.useRef();
  const email = React.useRef();
  const password = React.useRef();
  const confirmPassword = React.useRef();

  const doSubmit = async (e) => {
    e.preventDefault();
    // axios
    if (password.current.value !== confirmPassword.current.value) {
      alert("Passwords do not match");
      return;
    }
    const res = await axios.post(Path.User, {
      name: name.current.value,
      email: email.current.value,
      password: password.current.value,
    });

    if (res.status === 201) {
      alert("User created successfully");
    }
  };

  return (
    <div>
      <Box
        sx={{
          display: "flex",
          "& > :not(style)": {
            m: 1,
            width: 800,
            height: 450,
            margin: "auto",
          },
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
                <TextField label="UserName" margin="normal" inputRef={name} />
                <TextField label="email" margin="normal" inputRef={email} />
                <TextField
                  type="password"
                  label="Password"
                  inputRef={password}
                  margin="normal"
                />
                <TextField
                  type="password"
                  label="Confirm Password"
                  inputRef={confirmPassword}
                  margin="normal"
                />
              </Box>
              <Box sx={{ mx: "6rem", my: "2rem" }}>
                <Button type="submit" variant="contained" color="primary">
                  登録
                </Button>
              </Box>
            </form>
          </Grid>
        </Paper>
      </Box>
    </div>
  );
};

export default Signup;
