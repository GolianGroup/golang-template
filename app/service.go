package app

import (
	postgres_repositories "golang_template/internal/repositories/postgres"
	"golang_template/internal/services"
)

func (a *application) InitServices(repository postgres_repositories.UserRepository) services.UserService {
	return services.NewUserService(repository)
}
