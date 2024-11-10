package controllers

type Controllers interface {
	GetUserController() UserController
}

type controllers struct {
	userController UserController
}

func NewControllers() Controllers {
	userController := NewUserController()
	return &controllers{userController: userController}
}

func (c *controllers) GetUserController() UserController {
	return c.userController
}
