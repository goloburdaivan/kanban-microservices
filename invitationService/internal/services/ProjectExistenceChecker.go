package services

import (
	"fmt"
	"net/http"
)

type ProjectExistenceChecker interface {
	Get(projectId int) error
}

type ProjectCheckerAPI struct{}

func (p *ProjectCheckerAPI) Get(projectId int) error {
	result, err := http.Get(fmt.Sprintf("http://localhost:8082/projects/%d", projectId))
	if err != nil {
		return fmt.Errorf("Service unavailable")
	}

	defer result.Body.Close()

	if result.StatusCode != 200 {
		return fmt.Errorf("Project not found: %d %s", result.StatusCode, result.Status)
	}

	return nil
}

func NewProjectExistenceChecker() ProjectExistenceChecker {
	return &ProjectCheckerAPI{}
}
