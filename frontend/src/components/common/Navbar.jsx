import { useAuth } from "../../providers/authProvider";
import { Link, Outlet, useNavigate } from "react-router-dom";
import Button from "../Button";
import { useEffect } from "react";

const Navbar = () => {
  const { token, setToken } = useAuth();

  const navigate = useNavigate();

  const logout = () => {
    console.log("clearing token");
    setToken(null);
  };

  useEffect(() => {
    if (token !== null) {
      navigate("/");
    } else {
      navigate("/login");
    }
  }, [token]);

  if (token) {
    return (
      <div className="w-screen h-screen flex flex-col ">
        <div className="flex fixed w-screen backdrop-blur-xl grow ">
          <div className="p-4 w-10/12 ">
            <Link to={"/"}>
              <Button name={"Home"} />
            </Link>
          </div>
          <div className="p-4 w-2/12 ">
            <div className="flex justify-end">
              <div className="flex flex-col justify-center w-1/2"></div>
              <Button name={"Logout"} onClick={logout} />
            </div>
          </div>
        </div>
        <Outlet />
      </div>
    );
  } else {
    return (
      <div className="w-screen h-screen flex flex-col ">
        <div className="flex fixed w-screen backdrop-blur-xl grow ">
          <div className="p-4 w-10/12 ">
            <Link to={"/"}>
              <Button name={"Home"} />
            </Link>
          </div>
          <div className="p-4 w-2/12 ">
            <div className="flex justify-end">
              <div className="flex flex-col justify-center w-1/2"></div>
              <Link to={"/login"}>
                <Button name={"Login"} />
              </Link>
            </div>
          </div>
        </div>
        <Outlet />
      </div>
    );
  }
};
export default Navbar;
