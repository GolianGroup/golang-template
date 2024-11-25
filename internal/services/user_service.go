package services

import (
	repositories "golang_template/internal/repositories/postgres"
	"golang_template/internal/services/dto"

	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	Login(ctx *fiber.Ctx, user dto.User)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Login(ctx *fiber.Ctx, user dto.User) {
	s.repo.Get(ctx.Context(), user)
}
