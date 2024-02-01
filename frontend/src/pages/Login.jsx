import React from "react";
import axios from "axios";
import { useAuth } from "../providers/authProvider";
import Textfield from "../components/Textfield";
import Button from "../components/Button";
import withAuthRedirect from "../providers/withAuthRedirect";

const Login = () => {
  const { setToken } = useAuth();

  const login = async () => {
    try {
      const email = document.getElementById("Email").value;
      if (!email) {
        console.error("Please input email");
        return;
      }
      // Get jwt-token

      const response = await axios.post(
        "https://www.zhefuz.link:8443/api/v1/login",
        {
          email: document.getElementById("Email").value,
          password: document.getElementById("Password").value,
        },
      );
      const token = response.data["token"];
      setToken(token);
      console.log(token);
      localStorage.setItem("token", token);
    } catch (error) {}
  };
  return (
    <div className="flex flex-col grow">
      <div className="flex  py-28 justify-center  "></div>
      <div className="flex flex-col grow space-y-8 justify-center items-center h-2/3">
        <Textfield
          name={"Email"}
          type={"email"}
          placeholder={"example@example.com"}
        />
        <Textfield name={"Password"} type={"password"} placeholder={""} />
        <Button onClick={login} name={"Login"} />
      </div>
    </div>
  );
};

export default withAuthRedirect(Login);
