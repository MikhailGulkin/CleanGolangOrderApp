package interactors

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/di/factories/interactors/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/di/factories/interactors/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/di/factories/interactors/relay"
	"go.uber.org/fx"
)

var Module = fx.Options(
	product.Module,
	order.Module,

	relay.Module,
)
