import { RouterProvider, createBrowserRouter } from "react-router-dom";
import { useAuth } from "../Providers/AuthProvider";
import { ProtectedRoute } from "./ProtectedRoute";
import RegisterPage from "../Pages/RegisterPage";
import LoginPage from "../Pages/LoginPage";
import ProjectsBoard from "../Pages/HomePage";
import ProjectPage from "../Pages/ProjectPage";

const Routes = () => {
    const { token } = useAuth();

    const routesForPublic = [
        {
            path: "/service",
            element: <div>Service Page</div>,
        },
        {
            path: "/about-us",
            element: <div>About Us</div>,
        },
        {
            path: "/register",
            element: <RegisterPage />,
        },
        {
            path: "/login",
            element: <LoginPage />,
        },
    ];

    const routesForAuthenticatedOnly = [
        {
            path: "/",
            element: <ProtectedRoute />, // Wrap the component in ProtectedRoute
            children: [
                {
                    path: "/",
                    element: <ProjectsBoard />,
                },
                {
                    path: "/project",
                    element: <ProjectPage />
                },
                {
                    path: "/profile",
                    element: <div>User Profile</div>,
                },
                {
                    path: "/logout",
                    element: <div>Logout</div>,
                },
            ],
        },
    ];
    const router = createBrowserRouter([
        ...routesForPublic,
        ...(!token ? [] : []),
        ...routesForAuthenticatedOnly,
    ]);

    return <RouterProvider router={router} />;
};

export default Routes;