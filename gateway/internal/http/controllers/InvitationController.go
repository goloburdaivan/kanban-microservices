package controllers

import (
	"gateway/internal/http/requests"
	"gateway/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InvitationController struct {
	service *services.InvitationService
}

func (i *InvitationController) Invite(c *gin.Context) {
	var request requests.InviteRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	err := i.service.Invite(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Invite sent successfully",
	})
}

func NewInvitationController(service *services.InvitationService) *InvitationController {
	return &InvitationController{service: service}
}
