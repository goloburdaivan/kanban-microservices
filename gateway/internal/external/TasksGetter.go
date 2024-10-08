package external

import (
	"encoding/json"
	"fmt"
	"gateway/internal/dto"
	"io/ioutil"
	"net/http"
	"strconv"
)

type TasksGetter interface {
	GetAll(projectId int) ([]dto.ColumnDTO, error)
}

type TasksGetterAPI struct{}

func (t TasksGetterAPI) GetAll(projectId int) ([]dto.ColumnDTO, error) {
	result, err := http.Get("http://localhost:8084/columns?projectId=" + strconv.Itoa(projectId))
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	if result.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cannot fetch columns %d", result.StatusCode)
	}

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	var response struct {
		Columns []dto.ColumnDTO `json:"columns"`
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Columns, nil
}

func NewTasksGetterAPI() TasksGetter {
	return &TasksGetterAPI{}
}
