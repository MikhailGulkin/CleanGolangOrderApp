package routes

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api/controllers/routes/healthcheck"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api/controllers/routes/order"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api/controllers/routes/product"
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	productRoutes product.Routes,
	orderRoutes order.Routes,
	healthcheckRoutes healthcheck.Routes,
) Routes {
	return Routes{
		productRoutes,
		orderRoutes,
		healthcheckRoutes,
	}
}
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
