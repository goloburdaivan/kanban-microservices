package initialisation

import (
	"fmt"
	"go.uber.org/dig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"projectsService/internal/database"
	"projectsService/internal/database/models"
	"projectsService/internal/http/controllers"
	"projectsService/internal/queue"
	"projectsService/internal/queue/listeners"
	"projectsService/internal/repository"
	"projectsService/internal/services"
)

func InitServiceContainer(db *gorm.DB) *dig.Container {
	container := dig.New()

	container.Provide(func() *gorm.DB {
		return db
	})

	container.Provide(database.NewPaginator)
	container.Provide(queue.NewRabbitMQSender)
	container.Provide(repository.NewProjectRepository)
	container.Provide(repository.NewProjectUserRepository)
	container.Provide(repository.NewUserGetter)
	container.Provide(services.NewProjectsService)

	container.Provide(controllers.NewProjectController)
	container.Provide(listeners.NewUserAcceptedInviteListener)

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
		log.Fatal("Error connecting to database")
	}

	return db
}

func MigrateSchemas(db *gorm.DB) {
	err := db.AutoMigrate(&models.Project{}, &models.ProjectUser{})
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
}
