package cron

import (
	"context"
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/logger"
	brokerconfigurate "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/message_broker/broker_configurate"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/cron/engine"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/cron/handlers"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/cron/providers"
	"go.uber.org/fx"
)

var Module = fx.Options(
	providers.Module,
	fx.Invoke(Start),
)

func Start(
	lifecycle fx.Lifecycle,
	setup brokerconfigurate.Brokers, //nolint:all
	handlers handlers.Handlers, //nolint:all
	cron engine.CronController, //nolint:all
	logger logger.Logger,
) {
	setup.Setup()
	handlers.Setup()

	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					defer func() {
						if r := recover(); r != nil {
							logger.Info(fmt.Sprintf("Recovered when boot cron, r %s", r))
						}
					}()
					cron.Run()
				}()
				return nil
			},
			OnStop: func(context.Context) error {
				logger.Info("Stopping cron application")
				return nil
			},
		},
	)
}
