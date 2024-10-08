// repository/column_repository.go

package repository

import (
	"boardService/internal/models"
	"gorm.io/gorm"
)

type ColumnRepositoryInterface interface {
	Create(column *models.Column) error
	GetAll(projectId int) ([]models.Column, error)
}

type columnRepository struct {
	db *gorm.DB
}

func (c *columnRepository) Create(column *models.Column) error {
	return c.db.Create(column).Error
}

func (c *columnRepository) GetAll(projectId int) ([]models.Column, error) {
	var columns []models.Column
	err := c.db.Where("project_id = ?", projectId).Preload("Tasks").Find(&columns).Error
	return columns, err
}

func NewColumnRepository(db *gorm.DB) ColumnRepositoryInterface {
	return &columnRepository{db: db}
}
