package middleware

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/middleware"
	errorhandler "github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/middleware/errorHandler"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(errorhandler.NewErrorMiddleware),
	fx.Provide(middleware.NewMiddlewares),
)
