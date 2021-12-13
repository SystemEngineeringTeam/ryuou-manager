import { Box } from "@mui/system";
import axios from "axios";
import React from "react";
import Path from "../../.react.config";
import {
  Button,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
} from "@mui/material";

const SubmitList = () => {
  const [submits, setSubmits] = React.useState([]);

  const fetchSubmitList = async () => {
    const res = await axios.get(Path.Admin.Question);
    setSubmits(res.data);
    console.log(res.data);
  };

  React.useEffect(() => {
    fetchSubmitList();
  }, []);

  return (
    <Box>
      <h1>Submit List</h1>
      <Table>
        <TableHead>
          <TableRow>
            <TableCell>ID</TableCell>
            <TableCell>Title</TableCell>
            <TableCell>URL</TableCell>
            <TableCell>TeamID</TableCell>
            <TableCell>Pass</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {submits &&
            submits.map((submit) => {
              return (
                <TableRow>
                  <TableCell>{submit.id}</TableCell>
                  <TableCell>{submit.title}</TableCell>
                  <TableCell>{submit.url}</TableCell>
                  <TableCell>{submit.team_id}</TableCell>
                  <TableCell>
                    <Button variant="contained" color="primary">
                      Pass
                    </Button>
                  </TableCell>
                </TableRow>
              );
            })}
        </TableBody>
      </Table>
    </Box>
  );
};

export default SubmitList;
