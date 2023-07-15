package handlers

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/handlers/address"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/handlers/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/handlers/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/handlers/user"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	product.NewProductHandler,
	user.NewUserHandler,
	address.NewProductHandler,
	order.NewOrderHandler,
)
