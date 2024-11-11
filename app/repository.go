package app

import "golang_template/internal/repositories"

func (a *application) InitRepositories() repositories.UserRepository {
	return repositories.NewUserRepository(a.InitDatabase(), &a.ctx)
}
