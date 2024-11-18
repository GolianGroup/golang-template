package repositories

import (
	"context"
	"errors"
	"fmt"
	"golang_template/internal/database/postgres"
	"golang_template/internal/ent"
	"golang_template/internal/ent/user"
	"golang_template/internal/services/dto"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound = &RepositoryErr{
		Msg: "User with this credentials not found",
		Err: errors.New("record not found"),
	}
	ErrUserAlreadyExists = &RepositoryErr{
		Msg: "User with this credentials already exists",
		Err: errors.New("user already exists"),
	}
	ErrDatabase = &RepositoryErr{
		Msg: "Database error occured",
		Err: errors.New("database error"),
	}
	ErrInvalidCredentials = &RepositoryErr{
		Msg: "Invalid credentials provided",
		Err: errors.New("invalid credentials"),
	}
)

type UserRepository interface {
	Get(ctx context.Context, userDto dto.User) (*ent.User, error)
	Create(ctx context.Context, userData dto.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type userRepository struct {
	db postgres.Database
}

func test(err error) {
	if err == ErrUserNotFound {
		fmt.Println("user not found")
	}
}

func NewUserRepository(db postgres.Database) UserRepository {
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
		return nil, ErrDatabase
	}

	return userData, nil
}

func (r userRepository) Create(ctx context.Context, userData dto.User) error {
	exists, _ := r.db.EntClient().User.
		Query().
		Where(user.Username(userData.Username)).
		Exist(ctx)

	if exists {
		return ErrUserAlreadyExists
	}

	// Create new user
	_, err := r.db.EntClient().User.
		Create().
		SetUsername(userData.Username).
		SetPassword(userData.Password).
		Save(ctx)

	if ent.IsValidationError(err) {
		return ErrInvalidCredentials
	}
	if err != nil {
		return ErrDatabase
	}

	return nil
}

func (r userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := r.db.EntClient().User.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		return ErrDatabase
	}

	return nil
}

func (r userRepository) Update(ctx context.Context, id uuid.UUID) error {
	r.db.EntClient()
	err := r.db.EntClient().User.
		UpdateOneID(id).
		Exec(ctx)

	if ent.IsValidationError(err) {
		return ErrInvalidCredentials
	}

	if err != nil {
		return ErrDatabase
	}

	return nil
}
