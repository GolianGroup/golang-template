package services

import "master/internal/pkg/repositories"

type IUserService interface{}

type UserService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(repository repositories.IUserRepository) IUserService {
	return &UserService{
		userRepository: repository,
	}
}
