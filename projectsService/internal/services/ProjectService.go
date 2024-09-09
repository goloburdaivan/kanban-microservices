package services

import (
	"errors"
	"projectsService/internal/database/models"
	"projectsService/internal/dto"
	"projectsService/internal/enums"
	"projectsService/internal/http/requests"
	"projectsService/internal/repository"
)

type ProjectsService struct {
	repository             *repository.ProjectDatabaseRepository
	projectUsersRepository *repository.ProjectUserRepository
	userRepository         repository.UserGetter
}

func (p *ProjectsService) GetProjects(userID uint, page, limit int) ([]models.Project, dto.PaginationMetaDTO, error) {
	data, total, err := p.repository.GetProjects(userID, page, limit)
	if err != nil {
		return nil, dto.PaginationMetaDTO{}, errors.New("paginate error:" + err.Error())
	}

	meta := dto.PaginationMetaDTO{
		Total: total,
		Page:  page,
		Pages: (total + limit - 1) / limit,
	}

	return data, meta, nil
}

func (p *ProjectsService) GetPermissions(userID int, projectID int) (string, error) {
	return p.projectUsersRepository.GetPermission(userID, projectID)
}

func (p *ProjectsService) Create(data *requests.CreateProjectRequest) (*models.Project, error) {
	_, err := p.userRepository.GetUser(uint64(data.UserID))
	if err != nil {
		return nil, err
	}

	project := &models.Project{
		Name:        data.Name,
		Description: data.Description,
	}

	project, err = p.repository.Create(project)
	if err != nil {
		return nil, err
	}

	projectUser := &models.ProjectUser{
		ProjectID: int(project.ID),
		UserID:    int(data.UserID),
		Role:      enums.RoleOwner,
	}

	projectUser, err = p.projectUsersRepository.Create(projectUser)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (p *ProjectsService) Delete(id int) (*models.Project, error) {
	project, err := p.repository.FindProjectById(id)
	if err != nil {
		return nil, err
	}

	return p.repository.Delete(project)
}

func (p *ProjectsService) Find(id int) (*models.Project, error) {
	return p.repository.FindProjectById(id)
}

func NewProjectsService(
	repository *repository.ProjectDatabaseRepository,
	userRepository repository.UserGetter,
	projectUsersRepository *repository.ProjectUserRepository,
) *ProjectsService {
	return &ProjectsService{
		repository:             repository,
		userRepository:         userRepository,
		projectUsersRepository: projectUsersRepository,
	}
}
