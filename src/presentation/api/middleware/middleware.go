package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewErrorMiddleware),
	fx.Provide(NewMiddlewares),
)

type Middleware interface {
	Handle(c *gin.Context)
}

type Middlewares []gin.HandlerFunc

func NewMiddlewares(
	errorMiddleware ErrorMiddleware,
) Middlewares {
	return Middlewares{
		errorMiddleware.Handle,
	}
}
