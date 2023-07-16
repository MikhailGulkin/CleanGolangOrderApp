package main

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/logger"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/consumer"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/cron"
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
