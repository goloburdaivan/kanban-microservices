package requests

type ProjectCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	UserID      uint   `json:"user_id"`
}
