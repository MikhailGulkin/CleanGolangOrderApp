package handlers

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/handlers/address"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/controllers/handlers/product"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(product.NewProductHandler),
	fx.Provide(address.NewProductHandler),
)
