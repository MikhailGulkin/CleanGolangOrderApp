package middleware

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/middleware/errorHandler"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(errorhandler.NewErrorMiddleware),
	fx.Provide(NewMiddlewares),
)

type Middlewares []gin.HandlerFunc

func NewMiddlewares(
	errorMiddleware errorhandler.ErrorMiddleware,
) Middlewares {
	return Middlewares{
		errorMiddleware.Handle,
	}
}
