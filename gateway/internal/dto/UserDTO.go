package dto

type UserDTO struct {
	ID       int    `json:"ID"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
