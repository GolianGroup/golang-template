package app

import (
	postgres_repositories "golang_template/internal/repositories/postgres"
	"golang_template/internal/services"
)

func (a *application) InitServices(repository repositories.Repository) services.Service {
	return services.NewService(repository)
}
