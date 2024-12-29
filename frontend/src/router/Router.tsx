import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { HomePage } from '../pages/home/HomePage';
import { LoginPage } from '../pages/login/LoginPage';
import { PrivateRoute } from '../components/private-route/PrivateRoute';
import { DashboardPage } from '../pages/dashboard/DashboardPage';
import { UserProvider } from '../contexts/UserContext';
import { RegistrationPage } from '../pages/registration/RegistrationPage';
import { Navbar } from '../components/navbar/Navbar';
import { routes } from './Routes.ts';


export const Router = () => {
    return (
        <BrowserRouter>
            <UserProvider>
                <Navbar />
                <Routes>
                    {/* Public Routes */}
                    <Route path={routes.root} Component={HomePage} />
                    <Route path={routes.login} Component={LoginPage} />
                    <Route path={routes.register} Component={RegistrationPage} />

                    {/* Private Routes */}
                    <Route
                        path={routes.home}
                        element={
                            <PrivateRoute>
                                <HomePage />
                            </PrivateRoute>
                        }
                    />
                    < Route
                        path={routes.dashboard}
                        element={
                            <PrivateRoute>
                                <DashboardPage />
                            </PrivateRoute>
                        }
                    />
                </Routes>
            </UserProvider>
        </BrowserRouter>
    )
}
