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
import Path from "../.react.config";

const Ranking = () => {
  const [ranking, setRanking] = React.useState([]);

  const fetchRanking = async () => {
    const res = await axios.get(Path.Ranking);
    setRanking(res.data);
    console.log(res.data);
  };

  React.useEffect(() => {
    fetchRanking();
  }, []);

  return (
    <Box>
      <h1>Ranking</h1>
      <Table>
        <TableHead>
          <TableRow>
            <TableCell>Rank</TableCell>
            <TableCell>Name</TableCell>
            <TableCell>Score</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {ranking &&
            ranking.map((r) => {
              return (
                <TableRow key={r.id}>
                  <TableCell>{r.rank}</TableCell>
                  <TableCell>{r.name}</TableCell>
                  <TableCell>{r.score}</TableCell>
                </TableRow>
              );
            })}
        </TableBody>
      </Table>
    </Box>
  );
};

export default Ranking;
