package app

import (
	"golang_template/internal/database"
	"golang_template/internal/repositories"
)

func (a *application) InitRepositories(db database.Database) repositories.Repository {
	return repositories.NewRepository(db)
}
