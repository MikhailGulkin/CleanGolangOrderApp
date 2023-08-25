package providers

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api/prometheus"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/di/api/controllers"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/di/api/engine"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/di/api/middleware"
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
