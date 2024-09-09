package requests

type InvitationRequest struct {
	UserID    int `json:"user_id"`
	ProjectID int `json:"project_id"`
}
