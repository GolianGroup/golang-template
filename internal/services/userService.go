package services

import (
	"golang_template/internal/repositories"
	"golang_template/internal/services/dto"
)

type UserService interface {
	Login(user dto.User)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s userService) Login(user dto.User) {
	s.repo.Get(user)
}
