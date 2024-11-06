package app

import (
	"golang_template/internal/repositories"
	"golang_template/internal/services"
)

func (a *application) InitServices(repository repositories.UserRepository) services.UserService {
	return services.NewUserService(repository)
}
