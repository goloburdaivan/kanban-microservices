package controllers

import (
	"gateway/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BoardController struct {
	service *services.BoardService
}

func (b *BoardController) Index(c *gin.Context) {
	projectId, _ := strconv.Atoi(c.Param("id"))
	data, err := b.service.GetBoardColumns(projectId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	project, _ := c.Get("project")

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"project": project,
		"data":    data,
	})
}

func NewBoardController(service *services.BoardService) *BoardController {
	return &BoardController{service: service}
}
