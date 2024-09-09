package requests

type CreateProjectRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	UserID      uint   `json:"user_id" binding:"required"`
}
