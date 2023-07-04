package routes

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewProductRoutes),
	controllers.Module,
)

type Routes []Route
type Route interface {
	Setup()
}

func NewRoutes(
	productRoutes ProductRoutes,
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
