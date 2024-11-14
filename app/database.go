package app

import (
	"golang_template/internal/database"
	"golang_template/internal/logging"

	"go.uber.org/zap"
)

func (a *application) InitDatabase(logger logging.Logger) database.Database {
	db, err := database.NewDatabase(a.ctx, &a.config.DB)
	if err != nil {
		logger.Fatal("Failed to start database", zap.Error(err))

	}
	return db
}
