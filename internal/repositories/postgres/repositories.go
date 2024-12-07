package postgres_repositories

import (
	"context"
	"golang_template/internal/database/arango"
	"golang_template/internal/database/postgres"
	"log"
)

type PostgresRepository interface {
	UserRepository() UserRepository
	VideoRepository() VideoRepository
}

// var (
// ErrGlobal = errors.New("some global error")
// )

type repository struct {
	userRepository  UserRepository
	videoRepository VideoRepository
}

func NewRepository(db postgres.Database, arango arango.ArangoDB, db postgres.PostgresDatabase, ctx context.Context) Repository {
	userRepository := NewUserRepository(db)
	videoRepository, err := NewVideoRepository(arango, ctx)
	if err != nil {
		log.Panic("Failed to initialize video repository", err)
	}
	return &repository{
		userRepository:  userRepository,
		videoRepository: videoRepository,
	}
}

func (r *repository) UserRepository() UserRepository {
	return r.userRepository
}

func (r *repository) VideoRepository() VideoRepository {
	return r.videoRepository
}
