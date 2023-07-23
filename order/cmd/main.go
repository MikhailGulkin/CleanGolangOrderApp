package main

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/di"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/logger"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/consumer"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/cron"
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
