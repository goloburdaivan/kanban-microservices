package routes

import (
	"gateway/internal/http/controllers"
	"gateway/internal/middleware"
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
		api.GET("/projects/:id", middleware.UserInProjectMiddleware(), func(c *gin.Context) {
			c.JSON(200, gin.H{
				"code": 200,
			})
		})
	}
}
