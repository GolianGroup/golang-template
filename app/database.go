package app

import (
	"golang_template/internal/database/postgres"
	"log"
)

func (a *application) InitDatabase() postgres.Database {
	db, err := postgres.NewDatabase(a.ctx, &a.config.DB)
	if err != nil {
		log.Fatalf("failed to setup database: %s", err)
	}
	return db
}
