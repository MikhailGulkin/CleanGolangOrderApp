package routes

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/handlers"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/routes/address"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/routes/product"
	"go.uber.org/fx"
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	productRoutes product.Routes, //nolint:all
	addressRoutes address.Routes,
) Routes {
	return Routes{
		productRoutes,
		addressRoutes,
	}
}
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(product.NewProductRoutes),
	fx.Provide(address.NewAddressRoutes),
	fx.Provide(NewGroupRoutes),
	handlers.Module,
)
