package app

import (
	"golang_template/internal/database"
	"log"
)

func (a *application) InitDatabase() database.Database {
	db, err := database.NewDatabase(a.ctx, &a.config.DB)
	if err != nil {
		log.Fatalf("failed to setup database: %s", err)
	}
	return db
}
