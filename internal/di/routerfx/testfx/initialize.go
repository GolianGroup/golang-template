package testfx

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"master/internal/pkg/handlers/api"
	test "master/internal/pkg/handlers/api/test"
)

var Module = fx.Provide(
	fx.Annotate(provideTestRouter, fx.ResultTags(`group:"routers"`)),
)

func provideTestRouter(app *fiber.App) api.IRouter {
	return test.NewTestRouter()
}
