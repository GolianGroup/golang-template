package repositories

import (
	"context"
	"errors"
	"master/internal/pkg/models"
)

type IUserRepository interface {
	Get(ctx context.Context, id int) (string, error)
}

type UserRepository struct {
	client *models.Client
}

func NewUserRepository(client *models.Client) IUserRepository {
	return &UserRepository{client: client}
}

// dto
func (r UserRepository) Get(ctx context.Context, id int) (string, error) {
	return "salam", errors.New("not implemented")
}
