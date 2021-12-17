import { Button, Grid, Paper, TextField } from "@mui/material";
import { Box } from "@mui/system";
import axios from "axios";
import React from "react";
import { useNavigate } from "react-router-dom";
import Path from "../../.react.config";

const AnswerForm = (props) => {
  const answer = React.useRef();

  const navigate = useNavigate();

  const doSubmit = async (e) => {
    e.preventDefault();
    const res = await axios.post(Path.Question + "/1/" + props.questionID, {
      answer: answer.current.value,
    });
    navigate("/list");
  };

  return (
    <Box position="static">
      <Paper variant="outlined">
        <Grid container alignItems="center" justifyContent="center">
          <form onSubmit={doSubmit}>
            <TextField type="url" label="URL" inputRef={answer} />
            <Grid container alignItems="center" justifyContent="end">
              <Button variant="contained" color="primary" type="submit">
                Submit
              </Button>
            </Grid>
          </form>
        </Grid>
      </Paper>
    </Box>
  );
};

export default AnswerForm;
