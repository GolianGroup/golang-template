package app

import (
	"golang_template/internal/database/arango"
	"golang_template/internal/database/postgres"
	"golang_template/internal/repositories"
)

func (a *application) InitRepositories(db postgres.Database, arango arango.ArangoDB) repositories.Repository {
	return repositories.NewRepository(db, arango, a.ctx)
}
