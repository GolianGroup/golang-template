package app

import (
	"fmt"
	"golang_template/internal/config"
	"golang_template/internal/logging"

	"go.uber.org/zap"
)

func (a *application) InitLogger() (logging.Logger, error) {
	if a.config.Server.Mode == "production" {
		logger, err := config.NewLoggerConfig(&a.config.Logger).Build()
		if err != nil {
			return nil, fmt.Errorf("failed to initialize production logger: %w", err)
		}
		return logger, nil
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize development logger: %w", err)
	}
	return logging.NewZapLogger(logger), nil
}
