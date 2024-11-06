package app

import "golang_template/handler/controllers"

func (a *application) InitController() controllers.UserController {
	return controllers.NewUserController()
}
