package initialisation

import (
	"boardService/internal/http/controllers"
	"boardService/internal/models"
	"boardService/internal/repository"
	"boardService/internal/services"
	"fmt"
	"go.uber.org/dig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitServiceContainer(db *gorm.DB) *dig.Container {
	container := dig.New()

	container.Provide(func() *gorm.DB {
		return db
	})
	container.Provide(repository.NewColumnRepository)
	container.Provide(repository.NewTaskRepository)
	container.Provide(services.NewColumnService)
	container.Provide(services.NewTaskService)

	container.Provide(controllers.NewColumnController)
	container.Provide(controllers.NewTaskController)

	return container
}

func InitDatabase() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to models")
	}

	return db
}

func MigrateSchemas(db *gorm.DB) {
	db.AutoMigrate(&models.Column{}, &models.Task{}, &models.TaskComment{})
}
