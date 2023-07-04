package routes

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/handlers"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewProductRoutes),
	handlers.Module,
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	productRoutes ProductRoutes, //nolint:all
) Routes {
	return Routes{
		productRoutes,
	}
}
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
