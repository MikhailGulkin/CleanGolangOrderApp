package interactors

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/interactors/address"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/interactors/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/interactors/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/interactors/relay"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/interactors/user"
	"go.uber.org/fx"
)

var Module = fx.Options(
	product.Module,
	user.Module,
	address.Module,
	order.Module,

	relay.Module,
)
