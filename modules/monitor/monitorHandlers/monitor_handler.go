package monitorHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/peedans/GoEcommerce/config"
	"github.com/peedans/GoEcommerce/modules/entities"
	"github.com/peedans/GoEcommerce/modules/monitor"
)

type IMonitorHandler interface {
	HealthCheck(c *fiber.Ctx) error
}

type monitorHandler struct {
	cfg config.IConfig
}

func (m *monitorHandler) HealthCheck(c *fiber.Ctx) error {
	res := &monitor.Monitor{
		Name:    m.cfg.App().Name(),
		Version: m.cfg.App().Version(),
	}
	return entities.NewResponse(c).Success(fiber.StatusOK, res).Res()
}

func MonitorHandler(cfg config.IConfig) IMonitorHandler {
	return &monitorHandler{
		cfg: cfg,
	}
}
