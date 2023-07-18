package routes

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/controllers/routes/address"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/controllers/routes/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/controllers/routes/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/controllers/routes/user"
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	productRoutes product.Routes,
	addressRoutes address.Routes,
	userRoutes user.Routes,
	orderRoutes order.Routes,
) Routes {
	return Routes{
		productRoutes,
		addressRoutes,
		userRoutes,
		orderRoutes,
	}
}
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
