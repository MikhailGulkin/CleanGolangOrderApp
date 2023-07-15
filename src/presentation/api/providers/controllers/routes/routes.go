package routes

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/routes"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/routes/address"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/routes/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/routes/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/routes/user"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	routes.NewRoutes,
	routes.NewGroupRoutes,

	product.NewProductRoutes,
	user.NewUserRoutes,
	order.NewOrderRoutes,
	address.NewAddressRoutes,
)
