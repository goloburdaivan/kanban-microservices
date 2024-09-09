package repository

import (
	"fmt"
	"gorm.io/gorm"
	"projectsService/internal/database/models"
)

type ProjectUserRepository struct {
	db *gorm.DB
}

func (p *ProjectUserRepository) GetPermission(userId, projectId int) (string, error) {
	var record models.ProjectUser
	err := p.db.First(&record, "user_id = ? AND project_id = ?", userId, projectId).Error
	if record.ID == 0 {
		return "", fmt.Errorf("permissions Not Found")
	}

	if err != nil {
		return "", err
	}

	return record.Role, nil
}

func (p *ProjectUserRepository) Create(model *models.ProjectUser) (*models.ProjectUser, error) {
	return model, p.db.Create(model).Error
}

func NewProjectUserRepository(db *gorm.DB) *ProjectUserRepository {
	return &ProjectUserRepository{db: db}
}
