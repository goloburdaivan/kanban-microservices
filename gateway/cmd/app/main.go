package main

import (
	"gateway/internal/cache"
	"gateway/internal/config"
	"gateway/internal/http/controllers"
	"gateway/internal/initialisation"
	"gateway/internal/middleware"
	"gateway/internal/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config.LoadConfig()
	redisClient := initialisation.InitCache()
	redis := cache.NewRedisCache(redisClient)
	container := initialisation.InitServiceContainer(redisClient)
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.AuthRequired(redis))

	err := container.Invoke(func(
		userController *controllers.UserProfileController,
		invitationController *controllers.InvitationController,
	) {
		routes.ApiRoutes(r, userController, invitationController)
	})
	if err != nil {
		log.Println(err)
		return
	}

	r.Run()
}
