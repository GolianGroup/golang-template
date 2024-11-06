package app

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
	"golang_template/handler/routers"
	"log"
)

type Application interface {
	Setup()
}

type application struct {
	ctx context.Context
	//viper config
}

func NewApplication(ctx context.Context) Application {
	return &application{ctx: ctx}
}

// bootstrap

func (a *application) Setup() {
	//viper
	app := fx.New(
		fx.Provide(
			a.InitRouter,
			a.InitFramework,
			a.InitController,
			a.InitServices,
			a.InitRepositories,
		),
		fx.Invoke(func(app *fiber.App, router routers.Router) {
			router.Handle()
			log.Fatal(app.Listen(":3000"))
		}),
	)
	app.Run()
}
