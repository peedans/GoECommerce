package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/peedans/GoEcommerce/middlewares/middlewaresHandler"
	"github.com/peedans/GoEcommerce/middlewares/middlewaresRepositories"
	"github.com/peedans/GoEcommerce/middlewares/middlewaresUsecases"
	monitorHandlers "github.com/peedans/GoEcommerce/modules/monitor/monitorHandlers"
	"github.com/peedans/GoEcommerce/modules/users/usersHandlers"
	"github.com/peedans/GoEcommerce/modules/users/usersRepositories"
	"github.com/peedans/GoEcommerce/modules/users/usersUsecases"
)

type IModuleFactory interface {
	MonitorModule()
	UsersModule()
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

func (m *moduleFactory) UsersModule() {
	repository := usersRepositories.UsersRepository(m.s.db)
	usecase := usersUsecases.UsersUsecase(m.s.cfg, repository)
	handler := usersHandlers.UsersHandler(m.s.cfg, usecase)
	router := m.r.Group("users")
	router.Post("/signup", handler.SignUpCustomer)
	router.Post("/signin", handler.SingIn)
	router.Post("/refresh", handler.RefreshPassport)
	router.Post("/signout", handler.SignOut)
	router.Post("/signup-admin", handler.SignUpAdmin)

	router.Get("/:user_id", m.mid.JwtAuth(), m.mid.ParamCheck(), handler.GerUserProfile)
	router.Get("/admin/secret", m.mid.JwtAuth(), m.mid.Authorize(2), handler.GenerateAdminToken)
}
