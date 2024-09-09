package requests

type InviteRequest struct {
	UserID    int `json:"user_id"`
	ProjectID int `json:"project_id"`
}
