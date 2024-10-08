// src/components/KanbanBoard.js

import React, { useState, useEffect } from 'react';
import {
    Box,
    Typography,
    Avatar,
    IconButton,
    Grid,
    Card,
    CardContent,
    CardActions,
    Button,
    Paper,
    Skeleton,
    TextField,
    Dialog,
    DialogTitle,
    DialogContent,
    DialogActions,
    DialogContentText,
    Tooltip
} from '@mui/material';
import EditIcon from '@mui/icons-material/Edit';
import AddIcon from '@mui/icons-material/Add';
import SearchIcon from '@mui/icons-material/Search';
import WorkIcon from '@mui/icons-material/Work';
import ExitToAppIcon from '@mui/icons-material/ExitToApp';
import ContentCopyIcon from '@mui/icons-material/ContentCopy';
import axios from "axios";
import { useAuth } from "../Providers/AuthProvider";
import NewProjectModal from "../components/Modals/NewProjectModal";
import InviteUserModal from "../components/Modals/InviteUserModal"; // Импортируем новый компонент

function ProjectsBoard() {
    const { token, logout } = useAuth();
    const [loading, setLoading] = useState(true);
    const [user, setUser] = useState(null);
    const [projects, setProjects] = useState([]);
    const [filteredProjects, setFilteredProjects] = useState([]);
    const [searchTerm, setSearchTerm] = useState('');
    const [openNewProjectModal, setOpenNewProjectModal] = useState(false);
    const [openLogoutDialog, setOpenLogoutDialog] = useState(false);
    const [openInviteUserModal, setOpenInviteUserModal] = useState(false);
    const [currentProjectId, setCurrentProjectId] = useState(null); // Для хранения текущего ID проекта

    const handleProjectAdded = (newProject) => {
        setProjects(prevProjects => {
            const updatedProjects = [...prevProjects, newProject];
            setFilteredProjects(updatedProjects);
            return updatedProjects;
        });
    };

    useEffect(() => {
        axios.get("http://localhost:8080/api/user/profile", {
            headers: {
                "Authorization": `Bearer ${token}`
            }
        })
            .then(res => {
                setUser(res.data.data.user);
                const fetchedProjects = res.data.data.projects.paginated;
                setProjects(fetchedProjects);
                setFilteredProjects(fetchedProjects);
                setLoading(false);
            })
            .catch(err => {
                console.log(err);
            });
    }, [token]);

    const handleSearch = (event) => {
        const value = event.target.value;
        setSearchTerm(value);
        if (value) {
            setFilteredProjects(projects.filter(project =>
                project.name.toLowerCase().includes(value.toLowerCase())
            ));
        } else {
            setFilteredProjects(projects);
        }
    };

    const handleOpenNewProjectModal = () => setOpenNewProjectModal(true);
    const handleCloseNewProjectModal = () => setOpenNewProjectModal(false);
    const handleOpenLogoutDialog = () => setOpenLogoutDialog(true);
    const handleCloseLogoutDialog = () => setOpenLogoutDialog(false);
    const handleOpenInviteUserModal = (projectId) => {
        setCurrentProjectId(projectId); // Устанавливаем текущий проект
        setOpenInviteUserModal(true);
    };
    const handleCloseInviteUserModal = () => setOpenInviteUserModal(false);

    const handleLogout = () => {
        logout();
        handleCloseLogoutDialog();
    };

    const handleCopyToClipboard = (text) => {
        navigator.clipboard.writeText(text).then(() => {
        }).catch(err => {
            console.error('Failed to copy text: ', err);
        });
    };

    return (
        <Box sx={{ p: 4, bgcolor: '#f0f2f5', minHeight: '100vh' }}>
            <Paper elevation={3} sx={{ p: 3, mb: 4 }}>
                <Box display="flex" alignItems="center" justifyContent="space-between">
                    <Box display="flex" alignItems="center">
                        {loading ? (
                            <Skeleton variant="circular" width={64} height={64} sx={{ mr: 2 }} />
                        ) : (
                            <Avatar sx={{ width: 64, height: 64, mr: 2 }}>{user.username.substring(0, 2)}</Avatar>
                        )}
                        <Box>
                            <Typography variant="h4">Welcome back,</Typography>
                            {loading ? (
                                <Skeleton width={120} />
                            ) : (
                                <Box display="flex" alignItems="center">
                                    <Typography sx={{paddingRight: 3}} variant="h4" color="primary">{user.username}</Typography>
                                    <Typography variant="h4" color="primary">#{user.ID}</Typography>
                                    <Tooltip title="Copy ID to clipboard">
                                        <IconButton onClick={() => handleCopyToClipboard(user.ID)} sx={{ ml: 1 }}>
                                            <ContentCopyIcon />
                                        </IconButton>
                                    </Tooltip>
                                </Box>
                            )}
                        </Box>
                    </Box>
                    <Box>
                        <IconButton color="primary" aria-label="edit profile" disabled={loading}>
                            <EditIcon />
                        </IconButton>
                        <IconButton color="secondary" aria-label="logout" onClick={handleOpenLogoutDialog} disabled={loading}>
                            <ExitToAppIcon />
                        </IconButton>
                    </Box>
                </Box>
            </Paper>

            <Box display="flex" alignItems="center" justifyContent="space-between" mb={4}>
                <Box display="flex" alignItems="center" sx={{ bgcolor: 'white', p: 1, borderRadius: 2 }}>
                    <SearchIcon color="action" sx={{ mr: 1 }} />
                    <TextField
                        placeholder="Search projects..."
                        variant="outlined"
                        size="small"
                        value={searchTerm}
                        onChange={handleSearch}
                        sx={{ width: 300 }}
                    />
                </Box>
            </Box>
            <Grid container spacing={3}>
                {loading ? (
                    Array.from(new Array(3)).map((_, index) => (
                        <Grid item xs={12} sm={6} md={4} key={index}>
                            <Skeleton variant="rectangular" height={150} />
                        </Grid>
                    ))
                ) : (
                    filteredProjects.map((project, index) => (
                        <Grid item xs={12} sm={6} md={4} key={project.ID}>
                            <Card sx={{
                                height: '100%',
                                display: 'flex',
                                flexDirection: 'column',
                                justifyContent: 'space-between',
                                boxShadow: 3
                            }}>
                                <CardContent>
                                    <Box display="flex" alignItems="center" mb={2}>
                                        <WorkIcon />
                                        <Typography variant="h6" sx={{ ml: 1 }}>{project.name}</Typography>
                                    </Box>
                                    <Typography variant="body2" color="textSecondary">
                                        {project.description}
                                    </Typography>
                                </CardContent>
                                <CardActions>
                                    <Button size="small" color="primary">Open</Button>
                                    <Button size="small" color="secondary" onClick={() => handleOpenInviteUserModal(project.ID)}>
                                        Invite User
                                    </Button>
                                </CardActions>
                            </Card>
                        </Grid>
                    ))
                )}

                <Grid item xs={12} sm={6} md={4}>
                    <Card sx={{
                        height: '100%',
                        display: 'flex',
                        justifyContent: 'center',
                        alignItems: 'center',
                        boxShadow: 3,
                        border: '2px dashed #90caf9'
                    }}>
                        <IconButton color="primary" aria-label="add new project" sx={{ flexDirection: 'column' }} onClick={handleOpenNewProjectModal} disabled={loading}>
                            <AddIcon sx={{ fontSize: 50 }} />
                            <Typography variant="body1">New Project</Typography>
                        </IconButton>
                    </Card>
                </Grid>
            </Grid>
            <NewProjectModal
                open={openNewProjectModal}
                onClose={handleCloseNewProjectModal}
                onProjectAdded={handleProjectAdded}
            />
            <InviteUserModal
                open={openInviteUserModal}
                onClose={handleCloseInviteUserModal}
                projectId={currentProjectId}
            />
            <Dialog
                open={openLogoutDialog}
                onClose={handleCloseLogoutDialog}
                aria-labelledby="logout-dialog-title"
            >
                <DialogTitle id="logout-dialog-title">Logout</DialogTitle>
                <DialogContent>
                    <DialogContentText>
                        Are you sure you want to log out?
                    </DialogContentText>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleCloseLogoutDialog}>Cancel</Button>
                    <Button onClick={handleLogout} color="primary">Logout</Button>
                </DialogActions>
            </Dialog>
        </Box>
    );
}

export default ProjectsBoard;
