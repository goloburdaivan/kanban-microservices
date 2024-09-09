// src/components/Modals/InviteUserModal.js

import React, { useState } from 'react';
import {
    Dialog,
    DialogTitle,
    DialogContent,
    DialogActions,
    Button,
    TextField,
    Typography
} from '@mui/material';
import axios from 'axios';

function InviteUserModal({ open, onClose, projectId }) {
    const [userId, setUserId] = useState('');
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');

    const handleInvite = () => {
        const userIdNumber = Number(userId);
        if (isNaN(userIdNumber) || userId.trim() === '') {
            setError('Valid User ID is required');
            return;
        }
        setLoading(true);
        setError('');

        axios.post(`http://localhost:8080/api/invite`, {
            user_id: userIdNumber,
            project_id: projectId,
        })
            .then(() => {
                setUserId('');
                onClose();
            })
            .catch(err => {
                setError('Failed to invite user. Please try again.');
                console.error('Failed to invite user:', err);
            })
            .finally(() => {
                setLoading(false);
            });
    };

    return (
        <Dialog open={open} onClose={onClose}>
            <DialogTitle>Invite User to Project</DialogTitle>
            <DialogContent>
                <Typography variant="body1" mb={2}>
                    Enter the ID of the user you want to invite to this project.
                </Typography>
                <TextField
                    autoFocus
                    margin="dense"
                    label="User ID"
                    type="number" // Используем тип "number"
                    fullWidth
                    variant="outlined"
                    value={userId}
                    onChange={(e) => setUserId(e.target.value)}
                    disabled={loading}
                    error={!!error}
                    helperText={error}
                />
            </DialogContent>
            <DialogActions>
                <Button onClick={onClose} disabled={loading}>Cancel</Button>
                <Button onClick={handleInvite} color="primary" disabled={loading}>
                    Invite
                </Button>
            </DialogActions>
        </Dialog>
    );
}

export default InviteUserModal;
