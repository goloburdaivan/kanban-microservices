package models

import "gorm.io/gorm"

type ProjectUser struct {
	gorm.Model
	ProjectID int      `json:"project_id" binding:"required"`
	UserID    int      `json:"user_id" binding:"required"`
	Role      string   `json:"role"`
	Project   *Project `json:"project,omitempty"`
}
