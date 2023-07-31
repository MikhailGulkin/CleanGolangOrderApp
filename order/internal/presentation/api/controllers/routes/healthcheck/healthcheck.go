package healthcheck

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api/controllers/handlers/healthcheck"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api/engine"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api/middleware/logging"
)

type Routes struct {
	controller healthcheck.Handler
	logging.LoggerMiddleware
	engine.RequestHandler
}

func (r Routes) Setup() {
	r.RequestHandler.Gin.GET("healthcheck/", r.controller.GetStatus, r.LoggerMiddleware.Handle)
}

func NewHealthCheckRoutes(
	handler engine.RequestHandler,
	controller healthcheck.Handler,
	logging logging.LoggerMiddleware,
) Routes {
	return Routes{controller: controller, RequestHandler: handler, LoggerMiddleware: logging}
}
