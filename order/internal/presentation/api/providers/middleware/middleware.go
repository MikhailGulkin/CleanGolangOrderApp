package middleware

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/presentation/api/middleware"
	errorhandler "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/presentation/api/middleware/errorHandler"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/presentation/api/middleware/logging"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	errorhandler.NewErrorMiddleware,
	logging.NewLoggerMiddleware,
	middleware.NewMiddlewares,
)
