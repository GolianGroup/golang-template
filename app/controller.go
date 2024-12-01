package app

import (
	"golang_template/handler/controllers"
	"golang_template/internal/logging"
)

func (a *application) InitController(logger logging.Logger) controllers.Controllers {
	return controllers.NewControllers(logger)
}
