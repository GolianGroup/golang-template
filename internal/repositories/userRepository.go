package repositories

import "golang_template/internal/services/dto"

type UserRepository interface {
	Get(user dto.User)
}

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) UserRepository {
	return &userRepository{client: client}
}

// dto
func (r *userRepository) Get(user dto.User) {
	// Use ent client for queries
	// Example:
	// user, err := r.client.User.Query().Where(user.Name(user.Username)).First(ctx)
}
