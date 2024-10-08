package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"projectsService/internal/config"
	"projectsService/internal/http/controllers"
	"projectsService/internal/http/middleware"
	"projectsService/internal/initialisation"
	"projectsService/internal/queue"
	"projectsService/internal/queue/listeners"
)

func main() {
	config.LoadConfig()
	db := initialisation.InitDatabase()
	initialisation.MigrateSchemas(db)
	container := initialisation.InitServiceContainer(db)

	rabbitMQ := queue.NewRabbitMQ()

	err := container.Invoke(func(userAcceptedInviteListener *listeners.UserAcceptedInviteListener) {
		rabbitMQ.RegisterConsumer("user_accept", userAcceptedInviteListener.Listen)
	})
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	go func() {
		if err := rabbitMQ.StartConsuming(); err != nil {
			log.Fatalf("failed to start consuming: %v", err)
		}
	}()

	r := gin.Default()
	r.Use(middleware.ErrorHandler())

	container.Invoke(func(
		projectController *controllers.ProjectController,
	) {
		r.GET("/user/:userID/projects", projectController.GetProjects)
		r.GET("/projects/:id", projectController.GetProject)
		r.GET("/projects/:id/permissions", projectController.GetPermissions)
		r.POST("/projects", projectController.CreateProject)
		r.DELETE("/projects/:id", projectController.DeleteProject)
	})

	r.Run(":8082")
}
