package interactors

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/interactors/address"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/interactors/product"
	"go.uber.org/fx"
)

var Module = fx.Options(
	product.Module,
	address.Module,
)
