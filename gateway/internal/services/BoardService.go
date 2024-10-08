package services

import (
	"gateway/internal/dto"
	"gateway/internal/external"
)

type BoardService struct {
	tasksGetter external.TasksGetter
}

func (b *BoardService) GetBoardColumns(projectId int) ([]dto.ColumnDTO, error) {
	return b.tasksGetter.GetAll(projectId)
}

func NewBoardService(tasksGetter external.TasksGetter) *BoardService {
	return &BoardService{tasksGetter: tasksGetter}
}
