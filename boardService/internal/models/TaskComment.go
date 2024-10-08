package models

import "gorm.io/gorm"

type TaskComment struct {
	gorm.Model
	UserID  string `json:"user_id" binding:"required"`
	TaskID  string `json:"task_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}
