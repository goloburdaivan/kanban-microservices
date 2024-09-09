package controllers

import (
	"github.com/gin-gonic/gin"
	"invitationService/internal/http/requests"
	"invitationService/internal/services"
	"net/http"
)

type InvitationController struct {
	service *services.InvitationService
}

func (i *InvitationController) Invite(c *gin.Context) {
	var request requests.InvitationRequest
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

func (i *InvitationController) ConfirmInvitation(c *gin.Context) {
	token := c.DefaultQuery("token", "")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "token required",
		})
		return
	}

	err := i.service.Accept(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Your invitation has been confirmed",
	})
}

func NewInvitationController(service *services.InvitationService) *InvitationController {
	return &InvitationController{service: service}
}
