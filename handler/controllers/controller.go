package controllers

type IControllers interface {
	GetUserController() IUserController
}

type Controllers struct {
	userController IUserController
}

func NewControllers() IControllers {
	userController := NewUserController()
	return &Controllers{userController: userController}
}

func (c *Controllers) GetUserController() IUserController {
	return c.userController
}
