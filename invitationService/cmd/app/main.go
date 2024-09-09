package main

import (
	"github.com/gin-gonic/gin"
	"invitationService/internal/config"
	"invitationService/internal/http/controllers"
	"invitationService/internal/initialisation"
)

func main() {
	config.LoadConfig()
	db := initialisation.InitDatabase()
	initialisation.MigrateSchemas(db)
	container := initialisation.InitServiceContainer(db)

	r := gin.Default()
	container.Invoke(func(invitationController *controllers.InvitationController) {
		r.GET("/confirm", invitationController.ConfirmInvitation)
		r.POST("/invite", invitationController.Invite)
	})

	r.Run(":8083")
}
