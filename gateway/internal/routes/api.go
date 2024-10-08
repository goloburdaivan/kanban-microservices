package routes

import (
	"gateway/internal/cache"
	"gateway/internal/http/controllers"
	"gateway/internal/middleware"
	"github.com/gin-gonic/gin"
)

type Routing struct {
	userController       *controllers.UserProfileController
	invitationController *controllers.InvitationController
	boardController      *controllers.BoardController
	cache                *cache.RedisCache
}

func NewRouting(
	userController *controllers.UserProfileController,
	invitationController *controllers.InvitationController,
	boardController *controllers.BoardController,
	cache *cache.RedisCache,
) *Routing {
	return &Routing{
		userController:       userController,
		invitationController: invitationController,
		cache:                cache,
		boardController:      boardController,
	}
}

func (r *Routing) ApiRoutes(server *gin.Engine) {
	api := server.Group("/api")
	{
		api.GET("/user/profile", r.userController.Index)
		api.POST("/projects", r.userController.CreateProject)
		api.POST("/invite", r.invitationController.Invite)
		api.GET("/projects/:id", middleware.UserInProjectMiddleware(r.cache), r.boardController.Index)
	}
}
