package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	AssigneeID  int    `json:"assignee_id" binding:"required"`
	CreatorID   int    `json:"creator_id" binding:"required"`
	ColumnID    int    `json:"column_id" binding:"required"`
}
