package services

import "golang_template/internal/repositories"

type Service interface {
	UserService() UserService
	VideoService() VideoService
}

type service struct {
	userService  UserService
	videoService VideoService
}

func NewService(repo repositories.Repository) Service {
	userService := NewUserService(repo.UserRepository())
	videoService := NewVideoService(repo.VideoRepository())
	return &service{
		userService:  userService,
		videoService: videoService,
	}
}

func (s *service) UserService() UserService {
	return s.userService
}

func (s *service) VideoService() VideoService {
	return s.videoService
}

func (s *service) UserService() UserService {
	return s.userService
}
