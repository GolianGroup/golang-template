package app

import (
	"fmt"
	"golang_template/internal/config"
	"golang_template/internal/logging"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func (a *application) InitLogger() (logging.Logger, error) {
	if a.config.Server.Mode == "production" {
		// Config rotation
		ws := zapcore.AddSync(
			&lumberjack.Logger{
				Filename:   a.config.Logger.Rotation.Filename,
				MaxSize:    a.config.Logger.Rotation.MaxSize,
				MaxBackups: a.config.Logger.Rotation.MaxBackups,
				MaxAge:     a.config.Logger.Rotation.MaxAge,
			},
		)

		// Config encoder and syncer
		level, _ := zapcore.ParseLevel(a.config.Logger.Level)
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(config.NewLoggerEncoderConfig(&a.config.Logger.EncoderConfig)),
			ws,
			level,
		)
		logger := zap.New(core)
		return logging.NewZapLogger(logger), nil

	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize development logger: %w", err)
	}
	return logging.NewZapLogger(logger), nil
}
