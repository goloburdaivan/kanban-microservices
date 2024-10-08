package initialisation

import (
	"gateway/internal/cache"
	"gateway/internal/external"
	"gateway/internal/http/controllers"
	"gateway/internal/routes"
	"gateway/internal/services"
	redis "github.com/redis/go-redis/v9"
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
	container.Provide(external.NewTasksGetterAPI)
	container.Provide(services.NewUserService)
	container.Provide(services.NewProjectService)
	container.Provide(services.NewInvitationService)
	container.Provide(services.NewBoardService)
	container.Provide(controllers.NewUserProfileController)
	container.Provide(controllers.NewInvitationController)
	container.Provide(controllers.NewBoardController)
	container.Provide(routes.NewRouting)

	return container
}
