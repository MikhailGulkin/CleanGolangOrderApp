package handlers

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/handlers/address"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/handlers/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/handlers/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/handlers/user"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(product.NewProductHandler),
	fx.Provide(user.NewUserHandler),
	fx.Provide(address.NewProductHandler),
	fx.Provide(order.NewOrderHandler),
)
