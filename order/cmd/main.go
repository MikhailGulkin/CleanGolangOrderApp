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
	"os"
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
		fx.Options(fx.Invoke(func(diGraph fx.DotGraph) {
			err := os.WriteFile("./graph.txt", []byte(diGraph), 0644)
			if err != nil {
				return
			}
		},
		)),
	).Run()
}
