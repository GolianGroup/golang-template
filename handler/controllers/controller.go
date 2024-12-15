package controllers

import (
	"golang_template/internal/logging"
	"golang_template/internal/services"
)

type Controllers interface {
	UserController() UserController
	VideoController() VideoController
	RpcServiceController() RpcServiceController
}

type controllers struct {
	userController  UserController
	videoController VideoController
	rpcServiceController RpcServiceController
}


func NewControllers(s services.Service, logger logging.Logger) Controllers {
	userController := NewUserController(s.UserService(), logger)
	videoController := NewVideoController(s.VideoService())
	rpcServiceController := NewRpcServiceController(s.RpcServiceService(), logger)
	return &controllers{
		userController:  userController,
		videoController: videoController,
		rpcServiceController: rpcServiceController,
	}
}

func (c *controllers) UserController() UserController {
	return c.userController
}
	
func (c *controllers) VideoController() VideoController {
	return c.videoController
}

func (c *controllers) RpcServiceController() RpcServiceController {
	return c.rpcServiceController
}
