package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/peedans/GoEcommerce/middlewares/middlewaresHandler"
	"github.com/peedans/GoEcommerce/middlewares/middlewaresRepositories"
	"github.com/peedans/GoEcommerce/middlewares/middlewaresUsecases"
	monitorHandlers "github.com/peedans/GoEcommerce/modules/monitor/monitorHandlers"
)

type IModuleFactory interface {
	MonitorModule()
}

type moduleFactory struct {
	r   fiber.Router
	s   *server
	mid middlewaresHandler.IMiddlewaresHandler
}

func InitModule(r fiber.Router, s *server, mid middlewaresHandler.IMiddlewaresHandler) IModuleFactory {
	return &moduleFactory{
		r:   r,
		s:   s,
		mid: mid,
	}
}

func InitMiddlewares(s *server) middlewaresHandler.IMiddlewaresHandler {
	repository := middlewaresRepositories.MiddlewaresRepository(s.db)
	usecase := middlewaresUsecases.MiddlewaresRepository(repository)
	return middlewaresHandler.MiddlewaresRepository(s.cfg, usecase)
}

func (m *moduleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandler(m.s.cfg)

	m.r.Get("/", handler.HealthCheck)
}
