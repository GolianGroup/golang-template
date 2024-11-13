package app

import (
	"context"
	"golang_template/handler/routers"
	"golang_template/internal/database"
	"golang_template/internal/utils"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type Application interface {
	Setup()
}

type application struct {
	ctx    context.Context
	config *utils.Config
	//viper config
}

func NewApplication(ctx context.Context) Application {
	return &application{ctx: ctx, config: nil}
}

// bootstrap

func (a *application) Setup() {
	config, err := utils.SetupViper("config/config.yml")
	a.config = config

	if err != nil {
		log.Fatalf("failed to setup viper: %s", err.Error())
	}

	app := fx.New(
		fx.Provide(
			a.InitDatabase,
			a.InitRouter,
			a.InitFramework,
			a.InitController,
			a.InitServices,
			a.InitRepositories,
		),
		fx.Invoke(func(lc fx.Lifecycle, db database.Database) {
			lc.Append(fx.Hook{
				OnStop: func(ctx context.Context) error {
					log.Println(db.Close())
					return nil
				},
			})
		}),
		fx.Invoke(func(app *fiber.App, router routers.Router) {
			log.Fatal(app.Listen(a.config.Server.Host + ":" + a.config.Server.Port))
		}),
	)
	app.Run()
}
