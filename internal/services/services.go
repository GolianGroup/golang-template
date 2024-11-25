package services

import postgres_repositories "golang_template/internal/repositories/postgres"

type Service interface {
}

type service struct {
	userService UserService
}

func NewService(repo postgres_repositories.PostgresRepository) Service {
	userService := NewUserService(repo.UserRepository())
	return &service{userService: userService}
}
