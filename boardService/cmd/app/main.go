package main

import (
	"boardService/internal/config"
	"boardService/internal/http/controllers"
	"boardService/internal/initialisation"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	db := initialisation.InitDatabase()
	initialisation.MigrateSchemas(db)
	container := initialisation.InitServiceContainer(db)

	r := gin.Default()
	container.Invoke(func(
		columnsController *controllers.ColumnsController,
		tasksController *controllers.TasksController,
	) {
		r.GET("/columns", columnsController.GetAll)
		r.POST("/columns", columnsController.Create)
		r.POST("/tasks", tasksController.Create)
	})

	r.Run(":8084")
}
