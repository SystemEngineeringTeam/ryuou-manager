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

  const toPass = async (teamID, questionID) => {
    await axios.put(Path.Admin.Question + "/" + teamID + "/" + questionID);
    await fetchSubmitList();
  };

  return (
    <Box>
      <h1>Submit List</h1>
      <Table>
        <TableHead>
          <TableRow>
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
                <TableRow key={submit.id}>
                  <TableCell>{submit.title}</TableCell>
                  <TableCell>{submit.url}</TableCell>
                  <TableCell>{submit.team_id}</TableCell>
                  <TableCell>
                    <Button
                      variant="contained"
                      color="primary"
                      onClick={() => {
                        if (
                          window.confirm(
                            "正解にしますか？この操作は取り消しできません"
                          )
                        ) {
                          toPass(submit.team_id, submit.question_id);
                        }
                      }}
                    >
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
