import React from "react";
import { useParams } from "react-router-dom";
import {} from "@mui/material";

const QuestionInfo = () => {
  const { id } = useParams();

  console.log(id);

  return <div>{id}</div>;
};

export default QuestionInfo;
