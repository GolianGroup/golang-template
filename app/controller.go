package app

import (
	"golang_template/handler/controllers"
	"golang_template/internal/services"
)

func (a *application) InitController(service services.Service) controllers.Controllers {
	return controllers.NewControllers(service)
}
