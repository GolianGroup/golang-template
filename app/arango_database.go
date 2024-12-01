package app

import (
	"golang_template/internal/database/arango"
	"log"
)

func (a *application) InitArangoDB() arango.ArangoDB {
	db, err := arango.NewArangoDB(a.ctx, &a.config.ArangoDB)
	if err != nil {
		log.Fatalf("failed to setup arango database: %s", err)
	}
	return db
}
