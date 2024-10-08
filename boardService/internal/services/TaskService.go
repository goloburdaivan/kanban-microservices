package services

import (
	"boardService/internal/models"
	"boardService/internal/repository"
)

type TaskServiceInterface interface {
	Create(task *models.Task) error
}

type taskService struct {
	repository repository.TaskRepositoryInterface
}

func (t *taskService) Create(task *models.Task) error {
	return t.repository.Create(task)
}

func NewTaskService(repository repository.TaskRepositoryInterface) TaskServiceInterface {
	return &taskService{repository: repository}
}
