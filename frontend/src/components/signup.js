import React from "react";
import { Button, TextField } from "@mui/material";
import axios from "axios";

const Signup = () => {
  const name = React.useRef();
  const email = React.useRef();
  const password = React.useRef();
  const confirmPassword = React.useRef();

  const doSubmit = (e) => {
    e.preventDefault();
    // axios
    if (password.current.value !== confirmPassword.current.value) {
      alert("Passwords do not match");
      return;
    }
    console.log(password.current.value);
    console.log("submitted");
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
