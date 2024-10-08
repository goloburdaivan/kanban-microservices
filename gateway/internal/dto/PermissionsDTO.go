package dto

type PermissionsDTO struct {
	Role    string      `json:"role"`
	Project *ProjectDTO `json:"project"`
}
