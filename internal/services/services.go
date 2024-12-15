package services

import "golang_template/internal/repositories"

type Service interface {
	UserService() UserService
	VideoService() VideoService
	RpcServiceService() RpcServiceService
}

type service struct {
	userService  UserService
	videoService VideoService
	rpcServiceService RpcServiceService
}

func NewService(repo repositories.Repository) Service {
	userService := NewUserService(repo.UserRepository())
	videoService := NewVideoService(repo.VideoRepository())
	rpcServiceService := NewRpcServiceService()
	return &service{
		userService:  userService,
		videoService: videoService,
		rpcServiceService: rpcServiceService,
	}
}

func (s *service) UserService() UserService {
	return s.userService
}

func (s *service) VideoService() VideoService {
	return s.videoService
}

func (s *service) RpcServiceService() RpcServiceService {
	return s.rpcServiceService
}
