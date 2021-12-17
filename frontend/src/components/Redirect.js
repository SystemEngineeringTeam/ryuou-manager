import React from "react";
import { useNavigate } from "react-router-dom";

const Redirect = () => {
  const navigate = useNavigate();

  React.useEffect(() => {
    navigate("/list");
  }, []);

  return null;
};

export default Redirect;
