package models

import "gorm.io/gorm"

type Column struct {
	gorm.Model
	ProjectID int     `json:"project_id" binding:"required"`
	Title     string  `json:"title" binding:"required"`
	Order     int     `json:"order" binding:"required"`
	IsDone    bool    `json:"is_done"`
	Tasks     *[]Task `json:"tasks,omitempty"`
}
