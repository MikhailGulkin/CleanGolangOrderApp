package providers

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/presentation/api/prometheus"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/presentation/api/providers/controllers"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/presentation/api/providers/engine"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/presentation/api/providers/middleware"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"presentation.api",
	fx.Options(
		middleware.Module,
		engine.Module,
		controllers.Module,
		fx.Provide(
			prometheus.NewPrometheus,
		),
	),
)
