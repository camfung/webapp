import { UserType } from "../../enums/userType.enum";
import { useUser } from "../../hooks/useUser";
import { useState } from "react";
import { UserDashboard } from "../../components/dashboard/UserDashboard";
import { AdminDashboard } from "../../components/dashboard/AdminDashboard";
import { Box, Grid2 } from "@mui/material";

export const DashboardPage = () => {
    const { user } = useUser();
    const [isAdmin] = useState(() => {
        return user?.UserRoles.some(userRole => userRole.Role.ID === UserType.Admin);
    });

    return (
        <Box p={2} sx={{ display: 'flex', height: '100vh' }}>
            {isAdmin && user && <AdminDashboard user={user} />}
            {!isAdmin && user && <UserDashboard user={user} />}
        </Box >
    );
};
