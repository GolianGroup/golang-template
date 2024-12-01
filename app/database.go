package app

import (
	"golang_template/internal/database/arango"
	"golang_template/internal/database/postgres"
	"golang_template/internal/logging"

	"go.uber.org/zap"
)

func (a *application) InitDatabase(logger logging.Logger) postgres.Database {
	db, err := postgres.NewDatabase(a.ctx, &a.config.DB)
	if err != nil {
		logger.Fatal("Failed to start database", zap.Error(err))

	}
	return db
}

func (a *application) InitArangoDB() arango.ArangoDB {
	db, err := arango.NewArangoDB(a.ctx, &a.config.ArangoDB)
	if err != nil {
		log.Fatalf("failed to setup arango database: %s", err)
	}
	return db
}
