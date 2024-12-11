package rest

import (
	"maker-checker/service"
)

// StartServer initiate server
func StartServer(container *service.Container) *HttpServer {

	//Inject services instance from ServiceContainer
	userController := NewUserController(container.UserService)

	httpServer := NewHttpServer(
		container.GbeConfigService.GetConfig().Rest.Addr,
		userController,
	)

	go httpServer.Start()

	container.LoggerService.GetInstance().Info("rest server ok")
	return httpServer
}
