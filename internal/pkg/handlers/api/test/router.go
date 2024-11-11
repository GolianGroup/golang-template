package user

import (
	"github.com/gofiber/fiber/v2"
	"master/internal/pkg/handlers/api"
)

type testRouter struct {
}

func NewTestRouter() api.IRouter {
	return &testRouter{}
}

func (c testRouter) Handle(fiberRouter fiber.Router) {
	fiberRouter.Get("/user1", c.login)
	fiberRouter.Get("/user2", c.login)
}

func (c testRouter) login(ctx *fiber.Ctx) error {
	err := ctx.JSON("succeed")
	if err != nil {
		return err
	}
	return nil
	//make dto
}
