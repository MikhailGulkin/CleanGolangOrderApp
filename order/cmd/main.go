package main

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/di"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/logger"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/config"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/consumer"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/cron"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	fx.New(
		fx.WithLogger(func(logger logger.Logger) fxevent.Logger {
			return logger.GetFxLogger()
		}),
		di.Module,
		config.Module,
		api.Module,
		consumer.Module,
		cron.Module,
	).Run()
}
