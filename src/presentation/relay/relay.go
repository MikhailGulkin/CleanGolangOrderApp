package relay

import (
	"context"
	"fmt"
	brokerconfigurate "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/message_broker/broker_configurate"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/relay/cron"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		cron.NewCron,
		cron.NewController,
	),
	fx.Invoke(Start),
)

func Start(
	lifecycle fx.Lifecycle,
	setup brokerconfigurate.Brokers, //nolint:all
	cron cron.Controller,
) {
	setup.Setup()
	cron.Setup()

	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
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
