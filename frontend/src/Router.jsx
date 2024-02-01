import Navbar from "./components/common/Navbar";
import Login from "./pages/Login";
import Error from "./pages/Error";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Home from "./pages/Home";

const Router = () => {
  const routes = [
    {
      element: <Navbar />,
      children: [
        {
          path: "/",
          element: <Home />,
        },
        {
          path: "/login",
          element: <Login />,
        },
      ],
      errorElement: <Error />,
    },
  ];

  const router = createBrowserRouter(routes);

  return <RouterProvider router={router} />;
};
export default Router;
