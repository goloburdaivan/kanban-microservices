package controllers

import (
	"boardService/internal/models"
	"boardService/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ColumnController struct {
	service services.ColumnServiceInterface
}

func (cc *ColumnController) GetAll(c *gin.Context) {
	projectId, _ := strconv.Atoi(c.DefaultQuery("projectId", "0"))
	columns, err := cc.service.GetAll(projectId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"columns": columns,
	})
}

func (cc *ColumnController) Create(c *gin.Context) {
	var column models.Column
	if err := c.ShouldBindJSON(&column); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	err := cc.service.Create(&column)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Column created successfully",
	})
}

func NewColumnController(service services.ColumnServiceInterface) *ColumnController {
	return &ColumnController{service: service}
}
