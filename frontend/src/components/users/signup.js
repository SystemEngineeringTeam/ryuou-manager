import React from "react";
import { Button, TextField } from "@mui/material";
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
      <form onSubmit={doSubmit}>
        <TextField type="text" label="UserName" inputRef={name} />
        <TextField type="email" label="Email" inputRef={email} />
        <TextField type="password" label="Password" inputRef={password} />
        <TextField
          type="password"
          label="Confirm Password"
          inputRef={confirmPassword}
        />
        <Button type="submit" variant="contained" color="primary">
          登録
        </Button>
      </form>
    </div>
  );
};

export default Signup;
