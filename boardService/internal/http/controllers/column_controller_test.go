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

type mockColumnService struct {
	CreateFunc func(column *models.Column) error
	GetAllFunc func(projectId int) ([]models.Column, error)
}

func (m *mockColumnService) Create(column *models.Column) error {
	return m.CreateFunc(column)
}

func (m *mockColumnService) GetAll(projectId int) ([]models.Column, error) {
	return m.GetAllFunc(projectId)
}

func TestColumnController_GetAll(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		mockService := &mockColumnService{
			GetAllFunc: func(projectId int) ([]models.Column, error) {
				return []models.Column{
					{Title: "To Do", ProjectID: projectId},
					{Title: "In Progress", ProjectID: projectId},
				}, nil
			},
		}

		controller := NewColumnController(mockService)

		req, err := http.NewRequest(http.MethodGet, "/columns?projectId=1", nil)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req

		controller.GetAll(c)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, true, response["success"])
		assert.NotNil(t, response["columns"])
	})

	t.Run("Service Error", func(t *testing.T) {
		mockService := &mockColumnService{
			GetAllFunc: func(projectId int) ([]models.Column, error) {
				return nil, errors.New("database error")
			},
		}

		controller := NewColumnController(mockService)

		req, err := http.NewRequest(http.MethodGet, "/columns?projectId=1", nil)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req

		controller.GetAll(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, false, response["success"])
		assert.Equal(t, "database error", response["error"])
	})
}

func TestColumnController_Create(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		mockService := &mockColumnService{
			CreateFunc: func(column *models.Column) error {
				return nil
			},
		}

		controller := NewColumnController(mockService)

		column := models.Column{
			Title:     "Done",
			ProjectID: 1,
			Order:     3,
		}

		bodyBytes, err := json.Marshal(column)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/columns", bytes.NewBuffer(bodyBytes))
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
		assert.Equal(t, "Column created successfully", response["message"])
	})

	t.Run("Binding Error", func(t *testing.T) {
		mockService := &mockColumnService{}

		controller := NewColumnController(mockService)

		req, err := http.NewRequest(http.MethodPost, "/columns", bytes.NewBuffer([]byte("invalid json")))
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
}
