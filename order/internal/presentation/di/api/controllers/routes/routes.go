package routes

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api/controllers/routes"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api/controllers/routes/healthcheck"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api/controllers/routes/order"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api/controllers/routes/product"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	routes.NewRoutes,
	routes.NewGroupRoutes,

	product.NewProductRoutes,
	order.NewOrderRoutes,
	healthcheck.NewHealthCheckRoutes,
)
