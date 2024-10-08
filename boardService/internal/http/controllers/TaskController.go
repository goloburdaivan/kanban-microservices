package controllers

import (
	"boardService/internal/models"
	"boardService/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TaskController struct {
	service services.TaskServiceInterface
}

func (t *TaskController) Create(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	err := t.service.Create(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Task created successfully",
	})
}

func NewTaskController(service services.TaskServiceInterface) *TaskController {
	return &TaskController{
		service: service,
	}
}
