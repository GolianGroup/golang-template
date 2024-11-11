package userfx

import (
	"go.uber.org/fx"
	"master/internal/pkg/handlers/api"
	"master/internal/pkg/handlers/api/user"
	"master/internal/pkg/models"
	"master/internal/pkg/repositories"
	"master/internal/pkg/services"
)

var Module = fx.Provide(
	//provideUserDBRepository,
	provideUserRepository,
	fx.Annotate(provideUserRouter, fx.ResultTags(`group:"routers"`)),
	//fx.Annotate(provideUserRouter, fx.ResultTags(`name:"userRouter" group:"routers"`)),
	provideUserService,
)

//func provideUserDBRepository(db *models.Client) modelUser.DBRepository {
//	return modelUser.NewDBRepository(db)
//}

func provideUserRepository(client *models.Client) repositories.IUserRepository {
	return repositories.NewUserRepository(client)
}

func provideUserService(userRepo repositories.IUserRepository) services.IUserService {
	return services.NewUserService(userRepo)
}

func provideUserRouter(userSrv services.IUserService) api.IRouter {
	return user.NewUserRouter(userSrv)
}
