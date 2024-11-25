package app

import (
	"golang_template/internal/database/clickhouse"
	"golang_template/internal/database/postgres"
	"log"
)

func (a *application) InitPostgresDatabase() postgres.PostgresDatabase {
	db, err := postgres.NewPostgresDatabase(a.ctx, &a.config.Postgres)
	if err != nil {
		log.Fatalf("failed to setup postgres database: %s", err)
	}
	return db
}

func (a *application) InitClickhouseDatabase() clickhouse.ClickhouseDatabase {
	db, err := clickhouse.NewClickhouseDatabase(a.ctx, &a.config.Clickhouse)
	if err != nil {
		log.Fatalf("failed to setup clickhouse database: %s", err)
	}
	return db
}
