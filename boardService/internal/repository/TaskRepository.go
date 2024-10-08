package repository

import (
	"boardService/internal/models"
	"gorm.io/gorm"
)

type TaskRepositoryInterface interface {
	Create(task *models.Task) error
}

type taskRepository struct {
	db *gorm.DB
}

func (t *taskRepository) Create(task *models.Task) error {
	return t.db.Create(task).Error
}

func NewTaskRepository(db *gorm.DB) TaskRepositoryInterface {
	return &taskRepository{db: db}
}
