package di

import (
	"context"
	"go.uber.org/fx"
	"log"
	"master/internal/di/configfx"
	"master/internal/di/entfx"
	"master/internal/di/fiberfx"
	"master/internal/di/logfx"
	"master/internal/di/routerfx"
)

func Start() {
	app := fx.New(
		routerfx.Module,
		fiberfx.Module,
		entfx.Module,
		logfx.Module,
		configfx.Module,
		fx.Invoke(fiberfx.InitFiberApp),
	)

	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}
}
