package controllers

import (
	"golang_template/internal/logging"
	"golang_template/internal/services"
)

type Controllers interface {
	UserController() UserController
	VideoController() VideoController
}

type controllers struct {
	userController  UserController
	videoController VideoController
}


func NewControllers(s services.Service, logger logging.Logger) Controllers {
	userController := NewUserController(s.UserService(), logger)
	videoController := NewVideoController(s.VideoService())
	return &controllers{
		userController:  userController,
		videoController: videoController,
	}
}

func (c *controllers) UserController() UserController {
	return c.userController
}

func (c *controllers) VideoController() VideoController {
	return c.videoController
}
