import {
  Button,
  Grid,
  Paper,
  Select,
  MenuItem,
  FormControl,
  InputLabel,
} from "@mui/material";
import { Box } from "@mui/system";
import axios from "axios";
import React, { useEffect } from "react";
import Path from "../../.react.config";
import { useCookies } from "react-cookie";
import { useNavigate } from "react-router-dom";

const Frame = () => {
  const [cookie] = useCookies();

  const navigate = useNavigate();

  if (cookie.sessionID !== "admin") {
    navigate("/");
  }

  const [teamID, setTeamID] = React.useState("");
  const [userID, setUserID] = React.useState("");

  const [teamList, setTeamList] = React.useState([{ id: "", name: "" }]);
  const [userList, setUserList] = React.useState([
    { id: "", name: "", emal: "", password: "" },
  ]);

  useEffect(() => {
    axios.get(Path.Admin.Team).then((res) => {
      setTeamList(res.data);
    });

    axios.get(Path.User).then((res) => {
      setUserList(res.data);
    });
  }, []);

  const handleTeamChange = (event) => {
    setTeamID(event.target.value);
  };
  const handleUserIDChange = (event) => {
    setUserID(event.target.value);
  };

  const doSubmit = async (e) => {
    await axios.post(Path.Admin.Team + "/" + teamID + "/" + userID);
  };

  const doDelete = async (e) => {
    await axios.delete(Path.Admin.Team + "/member/" + userID);
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
              <h1>チーム管理</h1>
            </Box>

            {/* <Box
              sx={{
                margin: "auto",
              }}
            >
              <TextField
                placeholder="TeamID"
                type="TeamID"
                inputRef={teamID}
                margin="normal"
              />
            </Box> */}

            <Box
              sx={{
                marginLeft: "auto",
                marginRight: "auto",
                marginBottom: "2rem",
              }}
            >
              <Box sx={{ minWidth: 260 }}>
                <FormControl fullWidth>
                  <InputLabel id="demo-simple-select-label">TeamID</InputLabel>
                  <Select
                    labelId="demo-simple-select-label"
                    id="demo-simple-select"
                    value={teamID}
                    label="Age"
                    onChange={handleTeamChange}
                  >
                    {teamList.map((team) => (
                      <MenuItem value={team.id} key={team.id}>
                        {team.name}
                      </MenuItem>
                    ))}
                  </Select>
                </FormControl>
              </Box>
            </Box>

            <Box
              sx={{
                margin: "auto",
              }}
            >
              <Box sx={{ minWidth: 260 }}>
                <FormControl fullWidth>
                  <InputLabel id="demo-simple-select-label">UserID</InputLabel>
                  <Select
                    labelId="demo-simple-select-label"
                    id="demo-simple-select"
                    value={userID}
                    label="teams"
                    onChange={handleUserIDChange}
                  >
                    {userList.map((user) => (
                      <MenuItem value={user.id} key={user.id}>
                        {user.name}
                      </MenuItem>
                    ))}
                  </Select>
                </FormControl>
              </Box>
            </Box>
          </Box>

          <Grid item xs={6}>
            <Box sx={{ ml: "8rem", my: "2rem" }}>
              <Button type="submit" variant="contained" onClick={doSubmit}>
                チーム分け
              </Button>
            </Box>
          </Grid>
          <Grid item xs={6}>
            <Button
              type="submit"
              variant="contained"
              onClick={doDelete}
              color="error"
            >
              チーム除去
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Box>
  );
};

export default Frame;
