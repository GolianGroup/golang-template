package logfx

import (
	"go.uber.org/fx"
	"master/internal/pkg/log"
)

var Module = fx.Provide(initZap)

func initZap() *log.Logger {
	return log.NewLogger()
}
