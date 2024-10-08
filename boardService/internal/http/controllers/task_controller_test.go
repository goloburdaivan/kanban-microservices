package controllers

import (
	"boardService/internal/models"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockTaskService struct {
	CreateFunc func(task *models.Task) error
}

func (m *mockTaskService) Create(task *models.Task) error {
	return m.CreateFunc(task)
}

func TestTaskController_Create(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		mockService := &mockTaskService{
			CreateFunc: func(task *models.Task) error {
				return nil
			},
		}

		controller := NewTaskController(mockService)

		task := models.Task{
			Title:       "Test Task",
			Description: "This is a test task",
			AssigneeID:  1,
			CreatorID:   1,
			ColumnID:    1,
		}

		bodyBytes, err := json.Marshal(task)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(bodyBytes))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req

		controller.Create(c)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, true, response["success"])
		assert.Equal(t, "Task created successfully", response["message"])
	})

	t.Run("Binding Error", func(t *testing.T) {
		mockService := &mockTaskService{}

		controller := NewTaskController(mockService)

		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer([]byte("invalid json")))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req

		controller.Create(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, false, response["success"])
	})

	t.Run("Service Error", func(t *testing.T) {
		mockService := &mockTaskService{
			CreateFunc: func(task *models.Task) error {
				return errors.New("database error")
			},
		}

		controller := NewTaskController(mockService)

		task := models.Task{
			Title:       "Test Task",
			Description: "This is a test task",
			AssigneeID:  1,
			CreatorID:   1,
			ColumnID:    1,
		}

		bodyBytes, err := json.Marshal(task)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(bodyBytes))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req

		controller.Create(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, false, response["success"])
		assert.Equal(t, "database error", response["error"])
	})
}
