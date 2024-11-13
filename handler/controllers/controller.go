package controllers

type Controllers interface {
	UserController() UserController
}

type controllers struct {
	userController UserController
}

func NewControllers() Controllers {
	userController := NewUserController()
	return &controllers{userController: userController}
}

func (c *controllers) UserController() UserController {
	return c.userController
}
