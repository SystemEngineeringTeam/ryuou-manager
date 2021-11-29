import React from "react";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
} from "@mui/material";
import axios from "axios";
import Path from "./.react.config.js";

const QuestionList = () => {
  const fetchQuestions = async () => {
    const res = await axios.get(Path.Question + "/1");
    console.log(res);
  };

  React.useEffect(() => {
    fetchQuestions();
  }, []);

  const [questions, setQuestions] = React.useState();

  return (
    <Table>
      <TableHead>
        <TableRow>
          <TableCell>Title</TableCell>
          <TableCell>Description</TableCell>
        </TableRow>
      </TableHead>
      <TableBody>
        {questions &&
          questions.map((question) => (
            <TableRow key={question.id}>
              <TableCell>{question.title}</TableCell>
              <TableCell>{question.description}</TableCell>
            </TableRow>
          ))}
      </TableBody>
    </Table>
  );
};

export default QuestionList;
