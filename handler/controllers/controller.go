package controllers

import "golang_template/internal/logging"

type Controllers interface {
	UserController() UserController
}

type controllers struct {
	userController UserController
}

func NewControllers(logger logging.Logger) Controllers {
	userController := NewUserController(logger)
	return &controllers{userController: userController}
}

func (c *controllers) UserController() UserController {
	return c.userController
}
