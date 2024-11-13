package repositories

import (
	"context"
	"golang_template/internal/database"
	"golang_template/internal/ent"
	"golang_template/internal/ent/user"
	"golang_template/internal/services/dto"
)

type UserRepository interface {
	Get(userDto dto.User) (*ent.User, error)
	CreateUser(userData dto.User) error
}

type userRepository struct {
	db  database.Database
	ctx *context.Context
}

func NewUserRepository(ctx *context.Context, db database.Database) UserRepository {
	return &userRepository{
		db:  db,
		ctx: ctx,
	}
}

func (r userRepository) Get(userDto dto.User) (*ent.User, error) {
	userData, err := r.db.EntClient().User.
		Query().
		Where(user.Username(userDto.Username), user.Password(userDto.Password)).
		First(*r.ctx)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func (r userRepository) CreateUser(userData dto.User) error {
	_, err := r.db.EntClient().User.
		Create().
		SetUsername(userData.Username).
		SetPassword(userData.Password).
		Save(*r.ctx)
	if err != nil {
		return err
	}
	return nil
}
