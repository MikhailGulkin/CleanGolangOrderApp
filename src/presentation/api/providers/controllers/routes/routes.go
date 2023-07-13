package routes

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/routes"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/routes/address"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/routes/product"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(routes.NewRoutes),
	fx.Provide(product.NewProductRoutes),
	fx.Provide(address.NewAddressRoutes),
	fx.Provide(routes.NewGroupRoutes),
)
