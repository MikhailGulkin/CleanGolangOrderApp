package middleware

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/middleware"
	errorhandler "github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/middleware/errorHandler"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	errorhandler.NewErrorMiddleware,
	middleware.NewMiddlewares,
)
