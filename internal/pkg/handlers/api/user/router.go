package user

import (
	"github.com/gofiber/fiber/v2"
	"master/internal/pkg/handlers/api"
	"master/internal/pkg/services"
)

type userRouter struct {
	userSrv services.IUserService
}

func NewUserRouter(userSrv services.IUserService) api.IRouter {
	return &userRouter{userSrv: userSrv}
}

func (c userRouter) Handle(fiberRouter fiber.Router) {
	fiberRouter.Get("/user", c.login)
}

func (c userRouter) login(ctx *fiber.Ctx) error {
	err := ctx.JSON("succeed")
	if err != nil {
		return err
	}
	return nil
	//make dto
}
