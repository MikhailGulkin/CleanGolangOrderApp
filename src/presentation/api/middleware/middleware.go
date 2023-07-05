package middleware

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewErrorMiddleware),
	fx.Provide(NewMiddlewares),
)

type IMiddleware interface {
	Setup()
}

type Middlewares []IMiddleware

func NewMiddlewares(
	errorMiddleware ErrorMiddleware,
) Middlewares {
	return Middlewares{
		errorMiddleware,
	}
}

func (m Middlewares) Setup() {
	for _, middleware := range m {
		middleware.Setup()
	}
}
