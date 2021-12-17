import React from "react";
import {
  Button,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  Box,
} from "@mui/material";
import axios from "axios";
import Path from "../../.react.config";

const MembersList = () => {
  const [members, setMembers] = React.useState([]);

  const fetchMembersList = async () => {
    const res = await axios.get(Path.User);
    setMembers(res.data);
    console.log(res.data);
  };

  React.useEffect(() => {
    fetchMembersList();
  }, []);

  return (
    <Box>
      <h1>Members List</h1>
      <Table>
        <TableHead>
          <TableRow>
            <TableCell>Name</TableCell>
            <TableCell>TeamID</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {members &&
            members.map((member) => {
              return (
                <TableRow key={member.id}>
                  <TableCell>{member.name}</TableCell>
                  <TableCell>{member.teamID}</TableCell>
                </TableRow>
              );
            })}
        </TableBody>
      </Table>
    </Box>
  );
};

export default MembersList;
