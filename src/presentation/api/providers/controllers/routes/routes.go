package routes

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/routes"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/routes/address"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/routes/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/routes/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/routes/user"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(routes.NewRoutes),
	fx.Provide(routes.NewGroupRoutes),

	fx.Provide(product.NewProductRoutes),
	fx.Provide(user.NewUserRoutes),
	fx.Provide(order.NewOrderRoutes),
	fx.Provide(address.NewAddressRoutes),
)
