package repository

import (
	"errors"
	"gorm.io/gorm"
	"projectsService/internal/database"
	"projectsService/internal/database/models"
)

type ProjectDatabaseRepository struct {
	db        *gorm.DB
	paginator *database.Paginator
}

func (p *ProjectDatabaseRepository) GetProjects(userID uint, page, limit int) ([]models.Project, int, error) {
	var projects []models.Project
	var count int64

	if err := p.db.Model(&models.Project{}).
		Joins("JOIN project_users ON projects.id = project_users.project_id").
		Where("project_users.user_id = ?", userID).
		Count(&count).Error; err != nil {
		return nil, 0, err
	}

	result := p.paginator.Paginate(page, limit).Joins("LEFT JOIN project_users ON projects.id = project_users.project_id").
		Where("project_users.user_id = ?", userID).
		Find(&projects)

	return projects, int(count), result.Error
}

func (p *ProjectDatabaseRepository) Create(project *models.Project) (*models.Project, error) {
	result := p.db.Create(project)
	return project, result.Error
}

func (p *ProjectDatabaseRepository) Delete(project *models.Project) (*models.Project, error) {
	result := p.db.Delete(project)
	return project, result.Error
}

func (p *ProjectDatabaseRepository) FindProjectById(id int) (*models.Project, error) {
	project := &models.Project{}
	p.db.First(project, "id = ?", id)
	if project.ID == 0 {
		return project, errors.New("project not found")
	}

	return project, nil
}

func NewProjectRepository(db *gorm.DB, paginator *database.Paginator) *ProjectDatabaseRepository {
	return &ProjectDatabaseRepository{db: db, paginator: paginator}
}
