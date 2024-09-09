package initialisation

import (
	"fmt"
	"go.uber.org/dig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"invitationService/internal/http/controllers"
	"invitationService/internal/models"
	"invitationService/internal/queue"
	"invitationService/internal/repository"
	"invitationService/internal/services"
	"log"
	"os"
)

func InitServiceContainer(db *gorm.DB) *dig.Container {
	container := dig.New()

	container.Provide(func() *gorm.DB {
		return db
	})
	container.Provide(queue.NewRabbitMQSender)
	container.Provide(func(sender *queue.RabbitMQSender) func() {
		return func() {
			log.Println("Shutting down RabbitMQSender...")
			sender.Shutdown()
		}
	})

	container.Provide(repository.NewUserGetter)
	container.Provide(services.NewProjectExistenceChecker)
	container.Provide(repository.NewInvitationRepository)
	container.Provide(services.NewInvitationEmailService)
	container.Provide(services.NewInvitationService)
	container.Provide(controllers.NewInvitationController)
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
	db.AutoMigrate(&models.Invitation{})
}
