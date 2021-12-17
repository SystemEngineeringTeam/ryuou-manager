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

const Frame = () => {
  const [cookie, setCookie, removeCookie] = useCookies();

  const [teamName, setTeamName] = React.useState("");
  const [userName, setUserName] = React.useState("");

  const [teamList, setTeamList] = React.useState([{ id: "", name: "" }]);
  const [userList, setUserList] = React.useState([
    { id: "", name: "", emal: "", password: "" },
  ]);

  let teamID = null;
  let userID = null;

  useEffect(() => {
    axios.get(Path.Admin.Team).then((res) => {
      setTeamList(res.data);
    });

    axios.get(Path.User).then((res) => {
      setUserList(res.data);
    });
  }, []);

  const handleTeamChange = (event) => {
    setTeamName(event.target.value);
  };
  const handleUserNameChange = (event) => {
    setUserName(event.target.value);
  };

  const doSubmit = async (e) => {
    teamList.forEach((team) => {
      if (team.name === teamName) {
        teamID = team.id;
      }
    });

    userList.forEach((user) => {
      if (user.name === userName) {
        userID = user.id;
      }
    });

    const res = await axios.post(
      Path.Admin.Team + "/" + teamID + "/" + userID,
      {}
    );
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
              <h1>チーム分け</h1>
            </Box>

            {/* <Box
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
                  <InputLabel id="demo-simple-select-label">
                    TeamName
                  </InputLabel>
                  <Select
                    labelId="demo-simple-select-label"
                    id="demo-simple-select"
                    value={teamName}
                    label="Age"
                    onChange={handleTeamChange}
                  >
                    {teamList.map((team) => (
                      <MenuItem value={team.name} key={team.id}>
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
                  <InputLabel id="demo-simple-select-label">
                    UserName
                  </InputLabel>
                  <Select
                    labelId="demo-simple-select-label"
                    id="demo-simple-select"
                    value={userName}
                    label="teams"
                    onChange={handleUserNameChange}
                  >
                    {userList.map((user) => (
                      <MenuItem value={user.name} key={user.id}>
                        {user.name}
                      </MenuItem>
                    ))}
                  </Select>
                </FormControl>
              </Box>
            </Box>
          </Box>

          <Box sx={{ mx: "6rem", my: "2rem" }}>
            <Button type="submit" variant="contained" onClick={doSubmit}>
              決定
            </Button>
          </Box>
        </Grid>
      </Paper>
    </Box>
  );
};

export default Frame;
