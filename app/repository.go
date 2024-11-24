package app

import (
	"golang_template/internal/database/postgres"
	"golang_template/internal/repositories"
)

func (a *application) InitRepositories(db postgres.Database) repositories.Repository {
	return repositories.NewRepository(db)
}
