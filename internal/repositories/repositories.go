package repositories

import (
	"golang_template/internal/database"
)

type Repository interface {
	UserRepository() UserRepository
}

// var (
// ErrGlobal = errors.New("some global error")
// )

type repository struct {
	userRepository UserRepository
}

func NewRepository(db database.Database) Repository {
	userRepository := NewUserRepository(db)
	return &repository{userRepository: userRepository}
}

func (r *repository) UserRepository() UserRepository {
	return r.userRepository
}
