import React from "react";

import {
  AppBar,
  Box,
  Button,
  IconButton,
  Toolbar,
  Typography,
} from "@mui/material";
import { Menu } from "@mui/icons-material";
import { Link } from "react-router-dom";
import { useCookies } from "react-cookie";

const ApplicationBar = () => {
  const [cookies] = useCookies();

  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static" style={{ marginBottom: "20px" }}>
        <Toolbar>
          <Typography
            variant="h6"
            color="inherit"
            sx={{ flexGrow: 1 }}
            component={Link}
            to="/list"
          >
            竜王戦
          </Typography>
          {cookies.token !== "" ? (
            <div>
              <Button color="inherit" component={Link} to="/list">
                問題一覧
              </Button>
              <Button color="inherit" component={Link} to="/ranking">
                ランキング
              </Button>
              <Button
                color="inherit"
                href="https://github.com/SystemEngineeringTeam/ryuou-manager"
              >
                GitHub
              </Button>
            </div>
          ) : (
            <div>
              <Button color="inherit" component={Link} to="/signup">
                登録
              </Button>
              <Button color="inherit" component={Link} to="/login">
                ログイン
              </Button>
            </div>
          )}
        </Toolbar>
      </AppBar>
    </Box>
  );
};

export default ApplicationBar;
