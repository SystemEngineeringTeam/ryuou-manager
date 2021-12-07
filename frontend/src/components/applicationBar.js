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

const ApplicationBar = () => {
  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
          <IconButton
            edge="start"
            color="inherit"
            aria-label="menu"
            size="large"
          >
            <Menu />
          </IconButton>
          <Typography
            variant="h6"
            color="inherit"
            sx={{ flexGrow: 1 }}
            component={Link}
            to="/list"
          >
            竜王戦
          </Typography>
          <Button color="inherit" component={Link} to="/signup">
            登録
          </Button>
          <Button color="inherit" component={Link} to="/login">
            ログイン
          </Button>
        </Toolbar>
      </AppBar>
    </Box>
  );
};

export default ApplicationBar;
