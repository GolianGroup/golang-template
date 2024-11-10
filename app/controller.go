package app

import "golang_template/handler/controllers"

func (a *application) InitController() controllers.Controllers {
	return controllers.NewControllers()
}
