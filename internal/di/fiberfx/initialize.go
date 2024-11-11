package fiberfx

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"master/internal/pkg/config"
	"master/internal/pkg/handlers/api"
)

var Module = fx.Provide(initFiberEngine)

func InitFiberApp(app *fiber.App, config *config.Application, router api.Router, lc fx.Lifecycle) {
	for _, impl := range router.Routers {
		impl.Handle(app.Group(""))
		fmt.Println(impl)
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return app.Listen(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port))
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})
}

func initFiberEngine() *fiber.App {
	return fiber.New()
}
