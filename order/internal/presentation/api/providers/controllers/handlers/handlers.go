package handlers

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/presentation/api/controllers/handlers/healthcheck"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/presentation/api/controllers/handlers/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/presentation/api/controllers/handlers/product"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	product.NewProductHandler,
	order.NewOrderHandler,
	healthcheck.NewHealthCheckHandler,
)
