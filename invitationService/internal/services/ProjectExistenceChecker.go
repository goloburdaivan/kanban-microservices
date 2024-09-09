package external

import (
	"encoding/json"
	"fmt"
	"gateway/internal/dto"
	"io/ioutil"
	"net/http"
)

type ProjectGetter interface {
	GetAll(userId int) (*dto.ProjectsPaginatedDTO, error)
}

type ProjectGetterAPI struct{}

func (p *ProjectGetterAPI) GetAll(userId int) (*dto.ProjectsPaginatedDTO, error) {
	result, err := http.Get(fmt.Sprintf("http://localhost:8082/user/%d/projects", userId))
	if err != nil {
		return nil, err
	}

	defer result.Body.Close()

	jsonBytes, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	var project dto.ProjectsPaginatedDTO
	err = json.Unmarshal(jsonBytes, &project)

	if err != nil {
		return nil, err
	}

	return &project, nil
}

func NewProjectGetterAPI() ProjectGetter {
	return &ProjectGetterAPI{}
}
