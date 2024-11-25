package postgres_repositories

import (
	"golang_template/internal/database/postgres"
)

type PostgresRepository interface {
	UserRepository() UserRepository
}

// var (
// ErrGlobal = errors.New("some global error")
// )

type repository struct {
	userRepository UserRepository
}

func NewRepository(db postgres.PostgresDatabase) PostgresRepository {
	userRepository := NewUserRepository(db)
	return &repository{userRepository: userRepository}
}

func (r *repository) UserRepository() UserRepository {
	return r.userRepository
}
