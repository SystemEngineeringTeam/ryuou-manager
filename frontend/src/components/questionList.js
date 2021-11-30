import React from "react";
import {
  Button,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
} from "@mui/material";
import axios from "axios";
import Path from "./.react.config.js";

const QuestionList = () => {
  const [questions, setQuestions] = React.useState();

  const fetchQuestions = async (teamID) => {
    const res = await axios.get(Path.Question + "/1");
    setQuestions(res.data);
    console.log(questions);
  };

  const openQuestion = async (teamID, questionID) => {
    await axios.put(Path.Question + "/" + teamID + "/" + questionID);
  };

  const doOpenClick = (id) => {
    openQuestion(1, id);
    fetchQuestions();
  };

  React.useEffect(() => {
    fetchQuestions();
  }, [setQuestions]);

  return (
    <Table>
      <TableHead>
        <TableRow>
          <TableCell>Title</TableCell>
          <TableCell>Description</TableCell>
          <TableCell>OpenButton</TableCell>
        </TableRow>
      </TableHead>
      <TableBody>
        {questions &&
          questions.map((question) => (
            <TableRow key={question.id}>
              <TableCell>{question.title}</TableCell>
              <TableCell>{question.description}</TableCell>
              {question.is_opened ? (
                <TableCell>Opened</TableCell>
              ) : (
                <TableCell>
                  <Button
                    color="primary"
                    variant="contained"
                    onClick={() => {
                      doOpenClick(question.id);
                    }}
                  >
                    Open
                  </Button>
                </TableCell>
              )}
            </TableRow>
          ))}
      </TableBody>
    </Table>
  );
};

export default QuestionList;
