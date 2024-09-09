import { useAuth } from "../Providers/AuthProvider";
import { useEffect } from "react";
import {Navigate, Outlet} from "react-router-dom";

export const ProtectedRoute = () => {
    const { token, logout } = useAuth();

    useEffect(() => {
        const checkToken = () => {
            if (!token || isTokenExpired(token)) {
                logout();
            }
        };

        checkToken();
    }, [token, logout]);

    const isTokenExpired = (token) => {
        if (!token) return true;
        const payload = JSON.parse(atob(token.split('.')[1]));
        return payload.exp < Date.now() / 1000;
    };

    if (!token || isTokenExpired(token)) {
        return <Navigate to="/login" />;
    }

    return <Outlet />;
};
