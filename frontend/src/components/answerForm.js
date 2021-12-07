import { Button, Grid, Paper, TextField } from "@mui/material";
import { Box } from "@mui/system";
import React from "react";

const AnswerForm = () => {
  const doSubmit = (e) => {
    e.preventDefault();
  };
  return (
    <Box position="static">
      <Paper variant="outlined">
        <Grid container alignItems="center" justifyContent="center">
          <form onSubmit={doSubmit}>
            <TextField type="url" label="URL" />
            <Grid container alignItems="center" justifyContent="end">
              <Button variant="contained" color="primary">
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
