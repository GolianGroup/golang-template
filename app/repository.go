package app

import (
	"golang_template/internal/database/clickhouse"
	"golang_template/internal/database/postgres"
	clickhouse_repositories "golang_template/internal/repositories/clickhouse"
	postgres_repositories "golang_template/internal/repositories/postgres"
)

func (a *application) InitPostgresRepositories(db postgres.PostgresDatabase) postgres_repositories.PostgresRepository {
	return postgres_repositories.NewRepository(db)
}

func (a *application) InitClickhouseRepositories(db clickhouse.ClickhouseDatabase) clickhouse_repositories.ClickhouseRepository {
	return clickhouse_repositories.NewRepository(db)
}
