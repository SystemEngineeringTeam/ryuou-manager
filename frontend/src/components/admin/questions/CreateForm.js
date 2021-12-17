import React from "react";
import { Box, Button, Paper, TextField, Grid } from "@mui/material";
import Path from "../../.react.config";
import axios from "axios";
import { useNavigate } from "react-router-dom";

const CreateForm = () => {
  const title = React.useRef(null);
  const description = React.useRef(null);
  const score = React.useRef(null);

  const navigate = useNavigate();

  const doSubmit = async (e) => {
    e.preventDefault();
    const data = {
      title: title.current.value,
      description: description.current.value,
      score: Number(score.current.value),
    };

    const res = await axios.post(Path.Admin.Question, data);
    console.log(res);

    navigate("/admin");
  };

  return (
    <Box
      sx={{
        display: "flex",
        "& > :not(style)": {
          m: 1,
          width: "70%",
          height: "60%",
          margin: "auto",
        },
      }}
    >
      <Paper>
        <form onSubmit={doSubmit}>
          <Grid
            container
            direction="column"
            justifyContent="center"
            alignItems="center"
            style={{ padding: "20px", margin: "20px" }}
          >
            <TextField
              required
              type="text"
              inputRef={title}
              placeholder="Title"
              style={{ width: "90%", marginBottom: "20px" }}
            />
            <TextField
              required
              type="text"
              inputRef={description}
              placeholder="Description"
              multiline
              rows={5}
              style={{ width: "90%", marginBottom: "20px" }}
            />
            <TextField
              required
              type="number"
              inputRef={score}
              placeholder="Score"
              style={{ width: "90%", marginBottom: "20px" }}
            />
            <Button type="submit" variant="contained">
              Create
            </Button>
          </Grid>
        </form>
      </Paper>
    </Box>
  );
};

export default CreateForm;
