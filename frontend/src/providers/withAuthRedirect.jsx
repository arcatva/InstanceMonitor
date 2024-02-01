import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "./authProvider";

const withAuthRedirect = (WrappedComponent) => {
  return (props) => {
    const { token, setToken } = useAuth();
    const navigate = useNavigate();

    useEffect(() => {
      if (token) {
        navigate("/");
      } else {
        navigate("/login");
      }
    }, [token, navigate]);

    return <WrappedComponent {...props} />;
  };
};

export default withAuthRedirect;
