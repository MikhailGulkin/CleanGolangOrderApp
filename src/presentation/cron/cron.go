package cron

import (
	"context"
	"fmt"
	brokerconfigurate "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/message_broker/broker_configurate"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/cron/di"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/cron/engine"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/cron/handlers"
	"go.uber.org/fx"
)

var Module = fx.Options(
	di.Module,
	fx.Invoke(Start),
)

func Start(
	lifecycle fx.Lifecycle,
	setup brokerconfigurate.Brokers, //nolint:all
	handlers handlers.Handlers,
	cron engine.CronController,
) {
	setup.Setup()
	handlers.Setup()

	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go cron.Run()
				return nil
			},
			OnStop: func(context.Context) error {
				fmt.Println("Stopping application")
				return nil
			},
		},
	)
}
