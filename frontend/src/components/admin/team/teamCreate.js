import { Button, TextField, Grid, Paper } from "@mui/material";
import { Box } from "@mui/system";
import axios from "axios";
import React from "react";
import Path from "../../.react.config";

const Frame = () => {
  const teamName = React.useRef();

  const doSubmit = async (e) => {
    const res = await axios.post(Path.Admin.Team, {
      teamName: teamName.current.value,
    });
    console.log(res);
  };

  return (
    <Box
      sx={{
        m: 1,
        width: 500,
        height: 450,
        marginLeft: "auto",
        marginRight: "auto",
        marginTop: "5rem",
      }}
    >
      <Paper
        variant="outlined"
        style={{
          backgroundColor: "rgba(196,196,196,0.25)",
          borderWidth: "5px",
        }}
      >
        <Grid container justifyContent="center" alignItems="center">
          <Box
            sx={{
              display: "flex",
              flexDirection: "column",
              "& .MuiTextField-root": { width: "25ch" },
            }}
          >
            <Box
              sx={{
                margin: "auto",
              }}
            >
              <h1>チーム作成</h1>
            </Box>

            <Box
              sx={{
                margin: "auto",
              }}
            >
              <TextField
                placeholder="TeamName"
                type="TeamName"
                inputRef={teamName}
                margin="normal"
              />
            </Box>
          </Box>

          <Box sx={{ mx: "6rem", my: "2rem" }}>
            <Button type="submit" variant="contained" onClick={doSubmit}>
              作成
            </Button>
          </Box>
        </Grid>
      </Paper>
    </Box>
  );
};

export default Frame;
