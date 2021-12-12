import React from "react";
import { Box, Button, TextField } from "@mui/material";
import Path from "../../.react.config";
import axios from "axios";

const CreateForm = () => {
  const title = React.useRef(null);
  const description = React.useRef(null);
  const score = React.useRef(null);

  const doSubmit = async () => {
    await axios.post(Path.Admin.Questions);
  };

  console.log("hogeoge");

  return (
    <Box>
      <form onSubmit={doSubmit}>
        <TextField required type="text" inputRef={title} placeholder="Title" />
        <TextField
          required
          type="text"
          inputRef={description}
          placeholder="Description"
        />
        <TextField
          required
          type="number"
          inputRef={score}
          placeholder="Score"
        />
        <Button type="submit">Create</Button>
      </form>
    </Box>
  );
};

export default CreateForm;
