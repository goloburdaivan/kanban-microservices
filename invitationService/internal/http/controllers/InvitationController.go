package controllers

import (
	"github.com/gin-gonic/gin"
	"invitationService/internal/services"
	"net/http"
)

type InvitationConfirmationController struct {
	service *services.InvitationAcceptService
}

func (i *InvitationConfirmationController) ConfirmInvitation(c *gin.Context) {
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

func NewInvitationConfirmationController(service *services.InvitationAcceptService) *InvitationConfirmationController {
	return &InvitationConfirmationController{service: service}
}
