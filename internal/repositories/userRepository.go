package repositories

import "golang_template/internal/services/dto"

type UserRepository interface {
	Get(user dto.User)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

// dto
func (r userRepository) Get(user dto.User) {

}
