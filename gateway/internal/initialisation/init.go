package initialisation

import (
	"gateway/internal/cache"
	"gateway/internal/external"
	"gateway/internal/http/controllers"
	"gateway/internal/services"
	"github.com/redis/go-redis/v9"
	"go.uber.org/dig"
)

func InitCache() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func InitServiceContainer(r *redis.Client) *dig.Container {
	container := dig.New()

	container.Provide(func() *redis.Client {
		return r
	})

	container.Provide(cache.NewRedisCache)
	container.Provide(external.NewUserGetterAPI)
	container.Provide(external.NewProjectGetterAPI)
	container.Provide(services.NewUserService)
	container.Provide(services.NewProjectService)
	container.Provide(services.NewInvitationService)
	container.Provide(controllers.NewUserProfileController)
	container.Provide(controllers.NewInvitationController)

	return container
}
