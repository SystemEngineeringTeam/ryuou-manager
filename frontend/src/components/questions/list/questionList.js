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
import Path from "../../.react.config.js";
import { Link } from "react-router-dom";
import { useCookies } from "react-cookie";

const QuestionList = () => {
  const [questions, setQuestions] = React.useState([]);
  const [cookies, setCookies, removeCookies] = useCookies();

  const fetchQuestions = async () => {
    if (
      cookies.teamID === undefined ||
      cookies.teamID === null ||
      cookies.teamID === "" ||
      cookies.teamID === 0
    ) {
      setQuestions([
        {
          id: 0,
          title: "Please login",
          description: "Please login",
          is_opened: true,
        },
      ]);
      return;
    }
    const res = await axios.get(Path.Question + "/" + cookies.teamID);
    setQuestions(res.data);
  };

  const openQuestion = async (teamID, questionID) => {
    await axios.put(Path.Question + "/" + teamID + "/" + questionID);
  };

  const doOpenClick = async (id) => {
    await openQuestion(1, id);
    await fetchQuestions(1);
  };

  React.useEffect(() => {
    fetchQuestions(1);
  }, []);

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
                question.id === 0 ? (
                  <TableCell>
                    <Button
                      variant="contained"
                      color="primary"
                      component={Link}
                      to={"/login"}
                    >
                      ログイン
                    </Button>
                  </TableCell>
                ) : (
                  <TableCell>
                    <Button
                      variant="contained"
                      color="primary"
                      component={Link}
                      to={"/questionInfo/" + question.id}
                    >
                      詳細
                    </Button>
                  </TableCell>
                )
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
