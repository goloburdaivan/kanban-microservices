package main

import (
	"example.com/internal/config"
	"example.com/internal/http"
	"example.com/internal/initialisation"
	"example.com/internal/queue"
	"example.com/internal/queue/listeners"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.LoadConfig()
	db := initialisation.InitDatabase()
	initialisation.MigrateSchemas(db)
	container := initialisation.InitServiceContainer(db)

	rabbitMQ := queue.NewRabbitMQ()

	err := container.Invoke(func(userRegisteredListener *listeners.UserRegisteredListener) {
		rabbitMQ.RegisterConsumer("user_registered", userRegisteredListener.Listen)
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
	container.Invoke(func(controller *http.EmailController) {
		r.GET("/confirm", controller.Confirm)
	})

	r.Run(":8085")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	log.Println("Shutting down...")
	rabbitMQ.Shutdown()
	if err := container.Invoke(func(shutdownFunc func()) {
		shutdownFunc()
	}); err != nil {
		log.Fatalf("Error while invoking shutdown function: %v", err)
	}
}
