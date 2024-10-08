package services

import (
	"boardService/internal/models"
	"boardService/internal/repository"
)

type ColumnServiceInterface interface {
	Create(column *models.Column) error
	GetAll(projectId int) ([]models.Column, error)
}

type columnService struct {
	repository repository.ColumnRepositoryInterface
}

func (c *columnService) Create(column *models.Column) error {
	return c.repository.Create(column)
}

func (c *columnService) GetAll(projectId int) ([]models.Column, error) {
	return c.repository.GetAll(projectId)
}

func NewColumnService(repository repository.ColumnRepositoryInterface) ColumnServiceInterface {
	return &columnService{repository: repository}
}
