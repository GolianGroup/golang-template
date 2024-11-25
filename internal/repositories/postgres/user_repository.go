package postgres_repositories

import (
	"context"
	"errors"
	"fmt"
	"golang_template/internal/database/postgres"
	"golang_template/internal/ent"
	"golang_template/internal/ent/user"
	"golang_template/internal/services/dto"
	"log"
)

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrUserExists      = errors.New("username already exists")
	ErrInvalidPassword = errors.New("invalid password")
	ErrDatabaseError   = errors.New("database error")
)

type UserRepository interface {
	Get(ctx context.Context, userDto dto.User) (*ent.User, error)
	CreateUser(ctx context.Context, userData dto.User) error
}

type userRepository struct {
	db postgres.PostgresDatabase
}

func test(err error) {
	if err == ErrUserNotFound {
		fmt.Println("user not found")
	}
}

func NewUserRepository(db postgres.PostgresDatabase) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r userRepository) Get(ctx context.Context, userDto dto.User) (*ent.User, error) {
	userData, err := r.db.EntClient().User.
		Query().
		Where(user.Username(userDto.Username), user.Password(userDto.Password)).
		First(ctx)

	if ent.IsNotFound(err) {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("querying user: %w", err)
	}

	return userData, nil
}

func (r userRepository) CreateUser(ctx context.Context, userData dto.User) error {

	// Create new user
	_, err := r.db.EntClient().User.
		Create().
		SetUsername(userData.Username).
		SetPassword(userData.Password).
		Save(ctx)

	log.Println(err)

	if err != nil {
		return fmt.Errorf("creating user: %w", err)
	}

	return nil
}
