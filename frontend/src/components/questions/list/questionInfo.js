import React from "react";
import { useParams } from "react-router-dom";
import { Typography } from "@mui/material";
import axios from "axios";
import Path from "../../.react.config";
import { Box } from "@mui/system";
import { Grid } from "@mui/material";
import AnswerForm from "../answerForm/answerForm";

const QuestionInfo = () => {
  const { id } = useParams();

  const [question, setQuestion] = React.useState({});

  const fetchQuestion = async () => {
    const res = await axios.get(Path.Question + "/1/" + id);
    setQuestion(res.data);
  };

  React.useEffect(() => {
    fetchQuestion();
  }, []);

  return (
    <Box style={{ height: "100%" }}>
      <Grid container alignItems="center" justifyContent="center">
        <Grid item xs={8}>
          <Typography variant="h4">{question.title}</Typography>
          <hr />
          <Typography gutterBottom variant="h6">
            {question.description}
          </Typography>
          <AnswerForm questionID={id} />
        </Grid>
      </Grid>
    </Box>
  );
};

export default QuestionInfo;
