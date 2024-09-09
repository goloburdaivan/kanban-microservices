package controllers

import (
	"gateway/internal/http/requests"
	"gateway/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserProfileController struct {
	service        *services.UserService
	projectService *services.ProjectService
}

func (u *UserProfileController) Index(c *gin.Context) {
	userID := c.GetInt("userId")
	if userID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "user ID not found",
		})
		return
	}

	data, err := u.service.GetUserProfileData(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error getting user data. See details.",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (u *UserProfileController) CreateProject(c *gin.Context) {
	var request requests.ProjectCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	userID := c.GetInt("userId")
	if userID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "user ID not found",
		})
		return
	}

	request.UserID = uint(userID)
	result, err := u.projectService.CreateProject(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error creating project. See details.",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func NewUserProfileController(service *services.UserService) *UserProfileController {
	return &UserProfileController{service: service}
}
