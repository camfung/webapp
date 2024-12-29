import React, { useState } from 'react';
import { AppBar, Toolbar, Typography, Button, IconButton, Menu, MenuItem, Box } from '@mui/material';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import { useUser } from '../../hooks/useUser';
import { useNavigate } from 'react-router-dom';
import { routes } from '../../router/Routes';
import { useTranslation } from 'react-i18next';

export const Navbar: React.FC = () => {
    const { user, logout } = useUser();
    const navigate = useNavigate();
    const { t } = useTranslation();

    const [anchorEl, setAnchorEl] = useState<HTMLAnchorElement | null>(null);
    const open = Boolean(anchorEl);

    const handleProfileMenuOpen = (event: React.MouseEvent<HTMLAnchorElement>) => {
        setAnchorEl(event.currentTarget);
    };

    const handleMenuClose = () => {
        setAnchorEl(null);
    };

    const handleLogout = () => {
        handleMenuClose();
        logout();
    };

    return (
        <AppBar position="static" sx={{
            marginBottom: 2,
            backgroundColor: '#264653',
            boxShadow: '0 2px 4px rgba(0,0,0,0.08)'
        }}>
            <Toolbar>
                <Box sx={{ display: 'flex', alignItems: 'center', flexGrow: 1 }}>
                    <Typography
                        variant="h6"
                        sx={{
                            mr: 4,
                            cursor: 'pointer',
                            color: '#2A9D8F',
                            fontWeight: 600
                        }}
                        onClick={() => navigate(routes.home)}
                    >
                        {t('dictionary.AppName')}
                    </Typography>

                    {/* Navigation Links */}
                    <Box sx={{ display: 'flex', gap: 4 }}>
                        {['Text to Speech', 'Features', 'Pricing'].map((item) => (
                            <Button
                                key={item}
                                sx={{
                                    color: '#2A9D8F',
                                    textTransform: 'none',
                                    fontSize: '1rem',
                                    fontWeight: 500,
                                    '&:hover': {
                                        backgroundColor: 'transparent',
                                        color: '#3498db',
                                        transform: 'translateY(-4px)',
                                        transition: 'all 0.4s ease-in-out'
                                    }
                                }}
                            >
                                {item}
                            </Button>
                        ))}
                    </Box>
                </Box>

                {!user ? (
                    <Button
                        color="inherit"
                        onClick={() => navigate(routes.login)}
                        sx={{
                            color: '#2A9D8F',
                            '&:hover': {
                                backgroundColor: '#e9ecef'
                            }
                        }}
                    >
                        {t('button.login')}
                    </Button>
                ) : (
                    <>
                        <IconButton
                            color="inherit"
                            onClick={handleProfileMenuOpen}
                            component="span"
                            sx={{
                                color: '#2c3e50',
                                '&:hover': {
                                    backgroundColor: '#e9ecef'
                                }
                            }}
                        >
                            <AccountCircleIcon />
                        </IconButton>
                        <Menu
                            anchorEl={anchorEl}
                            open={open}
                            onClose={handleMenuClose}
                            anchorOrigin={{
                                vertical: 'top',
                                horizontal: 'right',
                            }}
                            transformOrigin={{
                                vertical: 'top',
                                horizontal: 'right',
                            }}
                            PaperProps={{
                                sx: {
                                    mt: 2,
                                    boxShadow: '0 4px 6px rgba(0,0,0,0.1)',
                                    '& .MuiMenuItem-root': {
                                        px: 3,
                                        py: 1.5,
                                        '&:hover': {
                                            backgroundColor: '#f8f9fa'
                                        }
                                    }
                                }
                            }}
                        >
                            <MenuItem onClick={() => { handleMenuClose(); navigate(routes.dashboard); }}>
                                {t('button.dashboard')}
                            </MenuItem>
                            <MenuItem onClick={handleLogout}>
                                {t('button.logout')}
                            </MenuItem>
                        </Menu>
                    </>
                )}
            </Toolbar>
        </AppBar>
    );
};
