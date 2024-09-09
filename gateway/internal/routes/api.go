package routes

import (
	"gateway/internal/http/controllers"
	"github.com/gin-gonic/gin"
)

func ApiRoutes(
	server *gin.Engine,
	userController *controllers.UserProfileController,
	invitationController *controllers.InvitationController,
) {
	api := server.Group("/api")
	{
		api.GET("/user/profile", userController.Index)
		api.POST("/projects", userController.CreateProject)
		api.POST("/invite", invitationController.Invite)
	}
}
