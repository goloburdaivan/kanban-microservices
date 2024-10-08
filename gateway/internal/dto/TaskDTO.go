package dto

type TaskDTO struct {
	ID          int    `json:"ID"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AssigneeID  int    `json:"assignee_id"`
	CreatorID   int    `json:"creator_id"`
}
