package routes

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/routes/address"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/routes/product"
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
