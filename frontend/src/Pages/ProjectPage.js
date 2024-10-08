import React, { useState, useEffect } from 'react';
import {
    Box,
    Typography,
    Paper,
    Button,
    Skeleton,
    TextField,
    Dialog,
    DialogActions,
    DialogContent,
    DialogTitle,
    IconButton,
    Chip,
    Avatar
} from '@mui/material';
import { Edit as EditIcon, Add as AddIcon } from '@mui/icons-material';
import { DragDropContext, Droppable, Draggable } from 'react-beautiful-dnd';
import axios from "axios";
import { useAuth } from "../Providers/AuthProvider";

function KanbanBoard() {
    const { token } = useAuth();
    const [loading, setLoading] = useState(true);
    const [columns, setColumns] = useState([]);
    const [openColumnModal, setOpenColumnModal] = useState(false);
    const [openTaskModal, setOpenTaskModal] = useState(false);
    const [openCreateTaskModal, setOpenCreateTaskModal] = useState(false);
    const [newColumnTitle, setNewColumnTitle] = useState('');
    const [project, setProject] = useState(null);
    const [selectedTask, setSelectedTask] = useState(null);
    const [newTask, setNewTask] = useState({
        title: '',
        description: '',
        status: '',
        priority: '',
        tags: '',
        assignee: ''
    });
    const [currentColumnId, setCurrentColumnId] = useState(null);

    useEffect(() => {
        axios.get("http://localhost:8080/api/projects/1", {
            headers: {
                "Authorization": `Bearer ${token}`
            }
        })
            .then(res => {
                setColumns(res.data.data);
                setProject(res.data.project.project);
                setLoading(false);
            })
            .catch(err => {
                console.log(err);
            });
    }, [token]);

    const onDragEnd = (result) => {
        const { source, destination } = result;

        if (!destination) return;
        if (source.droppableId === destination.droppableId && source.index === destination.index) return;

        const sourceColumnIndex = columns.findIndex(column => column.ID.toString() === source.droppableId);
        const destinationColumnIndex = columns.findIndex(column => column.ID.toString() === destination.droppableId);

        const sourceColumn = columns[sourceColumnIndex];
        const destinationColumn = columns[destinationColumnIndex];

        const sourceTasks = Array.from(sourceColumn.tasks);
        const [removedTask] = sourceTasks.splice(source.index, 1);

        if (sourceColumn === destinationColumn) {
            sourceTasks.splice(destination.index, 0, removedTask);
            const newColumns = [...columns];
            newColumns[sourceColumnIndex].tasks = sourceTasks;
            setColumns(newColumns);
        } else {
            const destinationTasks = Array.from(destinationColumn.tasks);
            destinationTasks.splice(destination.index, 0, removedTask);

            const newColumns = [...columns];
            newColumns[sourceColumnIndex].tasks = sourceTasks;
            newColumns[destinationColumnIndex].tasks = destinationTasks;

            setColumns(newColumns);
        }
    };

    const handleOpenColumnModal = () => setOpenColumnModal(true);
    const handleCloseColumnModal = () => setOpenColumnModal(false);

    const handleAddColumn = () => {
        const newColumn = {
            ID: columns.length + 1,
            title: newColumnTitle,
            tasks: []
        };
        setColumns([...columns, newColumn]);
        setNewColumnTitle('');
        handleCloseColumnModal();
    };

    const handleOpenTaskModal = (task) => {
        setSelectedTask({ ...task });
        setOpenTaskModal(true);
    };
    const handleCloseTaskModal = () => {
        setSelectedTask(null);
        setOpenTaskModal(false);
    };

    const handleOpenCreateTaskModal = (columnId) => {
        setCurrentColumnId(columnId);
        setNewTask({
            title: '',
            description: '',
            status: '',
            priority: '',
            tags: '',
            assignee: ''
        });
        setOpenCreateTaskModal(true);
    };
    const handleCloseCreateTaskModal = () => {
        setCurrentColumnId(null);
        setOpenCreateTaskModal(false);
    };

    const handleAddTask = () => {
        const columnIndex = columns.findIndex(column => column.ID === currentColumnId);
        const column = columns[columnIndex];
        const updatedTasks = [
            ...column.tasks,
            {
                ID: Date.now(), // Simple unique ID
                title: newTask.title,
                description: newTask.description,
                status: newTask.status,
                priority: newTask.priority,
                tags: newTask.tags.split(',').map(tag => tag.trim()),
                assignee: newTask.assignee
            }
        ];

        const newColumns = [...columns];
        newColumns[columnIndex].tasks = updatedTasks;
        setColumns(newColumns);
        handleCloseCreateTaskModal();
    };

    const handleSaveTask = () => {
        // Update the task in the columns state
        const newColumns = columns.map(column => {
            return {
                ...column,
                tasks: column.tasks.map(task => {
                    if (task.ID === selectedTask.ID) {
                        return selectedTask;
                    }
                    return task;
                })
            };
        });
        setColumns(newColumns);
        handleCloseTaskModal();
    };

    return (
        <Box sx={{ p: 4, bgcolor: '#f4f5f7', minHeight: '100vh' }}>
            <Typography variant="h4" gutterBottom>
                {project?.name}
            </Typography>

            <Button variant="contained" color="primary" onClick={handleOpenColumnModal} sx={{ mb: 2 }}>
                Add Column
            </Button>

            {loading ? (
                <Skeleton variant="rectangular" height={200} />
            ) : (
                <DragDropContext onDragEnd={onDragEnd}>
                    <Box
                        sx={{
                            display: 'flex',
                            overflowX: 'auto',
                            overflowY: 'hidden',
                            mt: 2,
                            alignItems: 'flex-start',
                            height: '80vh',
                            p: 1
                        }}
                    >
                        {columns.map((column) => (
                            <Box key={column.ID} sx={{ minWidth: '300px', mx: 1 }}>
                                <Paper elevation={4} sx={{
                                    p: 2,
                                    bgcolor: '#ebecf0',
                                    borderRadius: '8px',
                                    width: '300px',
                                    height: '100%',
                                    display: 'flex',
                                    flexDirection: 'column'
                                }}>
                                    <Box sx={{ display: 'flex', alignItems: 'center', mb: 1 }}>
                                        <Typography variant="h6" sx={{ fontWeight: 'bold', flexGrow: 1 }}>
                                            {column.title}
                                        </Typography>
                                        <IconButton
                                            size="small"
                                            onClick={() => handleOpenCreateTaskModal(column.ID)}
                                        >
                                            <AddIcon fontSize="small" />
                                        </IconButton>
                                    </Box>

                                    <Droppable droppableId={column.ID.toString()}>
                                        {(provided) => (
                                            <div
                                                {...provided.droppableProps}
                                                ref={provided.innerRef}
                                                style={{ minHeight: '100px', flexGrow: 1 }}
                                            >
                                                {column.tasks.map((task, index) => (
                                                    <Draggable key={task.ID} draggableId={task.ID.toString()} index={index}>
                                                        {(provided) => (
                                                            <Paper
                                                                ref={provided.innerRef}
                                                                {...provided.draggableProps}
                                                                {...provided.dragHandleProps}
                                                                sx={{
                                                                    mb: 2,
                                                                    p: 2,
                                                                    bgcolor: '#fff',
                                                                    borderRadius: '4px',
                                                                    boxShadow: '0px 1px 5px rgba(0, 0, 0, 0.1)',
                                                                    transition: 'box-shadow 0.2s ease',
                                                                    '&:hover': {
                                                                        boxShadow: '0px 4px 12px rgba(0, 0, 0, 0.15)'
                                                                    },
                                                                    position: 'relative'
                                                                }}
                                                            >
                                                                <IconButton
                                                                    size="small"
                                                                    sx={{ position: 'absolute', top: 4, right: 4 }}
                                                                    onClick={() => handleOpenTaskModal(task)}
                                                                >
                                                                    <EditIcon fontSize="small" />
                                                                </IconButton>
                                                                <Typography variant="body1" sx={{ fontWeight: 'bold' }}>
                                                                    {task.title}
                                                                </Typography>
                                                                <Typography variant="body2" color="textSecondary" sx={{ mb: 1 }}>
                                                                    {task.description}
                                                                </Typography>
                                                                <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5, mb: 1 }}>
                                                                    {task.tags?.map((tag, idx) => (
                                                                        <Chip key={idx} label={tag} size="small" />
                                                                    ))}
                                                                </Box>
                                                                {/* Assignee Stub */}
                                                                {task.assignee && (
                                                                    <Box sx={{ display: 'flex', alignItems: 'center', mt: 1 }}>
                                                                        <Avatar sx={{ width: 24, height: 24, mr: 1 }}>
                                                                            {task.assignee.charAt(0).toUpperCase()}
                                                                        </Avatar>
                                                                        <Typography variant="body2">
                                                                            {task.assignee}
                                                                        </Typography>
                                                                    </Box>
                                                                )}
                                                            </Paper>
                                                        )}
                                                    </Draggable>
                                                ))}
                                                {provided.placeholder}
                                            </div>
                                        )}
                                    </Droppable>
                                </Paper>
                            </Box>
                        ))}
                    </Box>
                </DragDropContext>
            )}

            {/* Modal for adding a new column */}
            <Dialog open={openColumnModal} onClose={handleCloseColumnModal}>
                <DialogTitle>Add New Column</DialogTitle>
                <DialogContent>
                    <TextField
                        autoFocus
                        margin="dense"
                        label="Column Title"
                        fullWidth
                        variant="outlined"
                        value={newColumnTitle}
                        onChange={(e) => setNewColumnTitle(e.target.value)}
                    />
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleCloseColumnModal} color="secondary">
                        Cancel
                    </Button>
                    <Button onClick={handleAddColumn} color="primary">
                        Add
                    </Button>
                </DialogActions>
            </Dialog>

            {/* Modal for editing a task */}
            <Dialog open={openTaskModal} onClose={handleCloseTaskModal} maxWidth="sm" fullWidth>
                <DialogTitle>Edit Task</DialogTitle>
                <DialogContent>
                    {selectedTask && (
                        <>
                            <TextField
                                margin="dense"
                                label="Title"
                                fullWidth
                                variant="outlined"
                                value={selectedTask.title}
                                onChange={(e) => {
                                    setSelectedTask({ ...selectedTask, title: e.target.value });
                                }}
                            />
                            <TextField
                                margin="dense"
                                label="Description"
                                fullWidth
                                multiline
                                rows={4}
                                variant="outlined"
                                value={selectedTask.description}
                                onChange={(e) => {
                                    setSelectedTask({ ...selectedTask, description: e.target.value });
                                }}
                                sx={{ mt: 2 }}
                            />
                            <TextField
                                margin="dense"
                                label="Status"
                                fullWidth
                                variant="outlined"
                                value={selectedTask.status || ''}
                                onChange={(e) => {
                                    setSelectedTask({ ...selectedTask, status: e.target.value });
                                }}
                                sx={{ mt: 2 }}
                            />
                            <TextField
                                margin="dense"
                                label="Priority"
                                fullWidth
                                variant="outlined"
                                value={selectedTask.priority || ''}
                                onChange={(e) => {
                                    setSelectedTask({ ...selectedTask, priority: e.target.value });
                                }}
                                sx={{ mt: 2 }}
                            />
                            <TextField
                                margin="dense"
                                label="Assignee"
                                fullWidth
                                variant="outlined"
                                value={selectedTask.assignee || ''}
                                onChange={(e) => {
                                    setSelectedTask({ ...selectedTask, assignee: e.target.value });
                                }}
                                sx={{ mt: 2 }}
                            />
                            <TextField
                                margin="dense"
                                label="Tags (comma-separated)"
                                fullWidth
                                variant="outlined"
                                value={selectedTask.tags?.join(', ') || ''}
                                onChange={(e) => {
                                    const tags = e.target.value.split(',').map(tag => tag.trim());
                                    setSelectedTask({ ...selectedTask, tags });
                                }}
                                sx={{ mt: 2 }}
                            />
                            <Typography variant="h6" sx={{ mt: 3 }}>
                                Comments
                            </Typography>
                            {/* Stub for comments */}
                            <Box sx={{ mt: 1 }}>
                                <Typography variant="body2" color="textSecondary">
                                    Comments will be displayed here...
                                </Typography>
                            </Box>
                        </>
                    )}
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleCloseTaskModal} color="secondary">
                        Cancel
                    </Button>
                    <Button
                        onClick={handleSaveTask}
                        color="primary"
                    >
                        Save
                    </Button>
                </DialogActions>
            </Dialog>

            {/* Modal for creating a new task */}
            <Dialog open={openCreateTaskModal} onClose={handleCloseCreateTaskModal} maxWidth="sm" fullWidth>
                <DialogTitle>Add New Task</DialogTitle>
                <DialogContent>
                    <TextField
                        autoFocus
                        margin="dense"
                        label="Title"
                        fullWidth
                        variant="outlined"
                        value={newTask.title}
                        onChange={(e) => setNewTask({ ...newTask, title: e.target.value })}
                    />
                    <TextField
                        margin="dense"
                        label="Description"
                        fullWidth
                        multiline
                        rows={4}
                        variant="outlined"
                        value={newTask.description}
                        onChange={(e) => setNewTask({ ...newTask, description: e.target.value })}
                        sx={{ mt: 2 }}
                    />
                    <TextField
                        margin="dense"
                        label="Status"
                        fullWidth
                        variant="outlined"
                        value={newTask.status}
                        onChange={(e) => setNewTask({ ...newTask, status: e.target.value })}
                        sx={{ mt: 2 }}
                    />
                    <TextField
                        margin="dense"
                        label="Priority"
                        fullWidth
                        variant="outlined"
                        value={newTask.priority}
                        onChange={(e) => setNewTask({ ...newTask, priority: e.target.value })}
                        sx={{ mt: 2 }}
                    />
                    <TextField
                        margin="dense"
                        label="Assignee"
                        fullWidth
                        variant="outlined"
                        value={newTask.assignee}
                        onChange={(e) => setNewTask({ ...newTask, assignee: e.target.value })}
                        sx={{ mt: 2 }}
                    />
                    <TextField
                        margin="dense"
                        label="Tags (comma-separated)"
                        fullWidth
                        variant="outlined"
                        value={newTask.tags}
                        onChange={(e) => setNewTask({ ...newTask, tags: e.target.value })}
                        sx={{ mt: 2 }}
                    />
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleCloseCreateTaskModal} color="secondary">
                        Cancel
                    </Button>
                    <Button onClick={handleAddTask} color="primary">
                        Add Task
                    </Button>
                </DialogActions>
            </Dialog>
        </Box>
    );
}

export default KanbanBoard;
