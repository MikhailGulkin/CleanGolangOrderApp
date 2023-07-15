package relay

import (
	"context"
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/relay/interfaces/interactors"
	brokerconfigurate "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/message_broker/broker_configurate"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Invoke(Start),
)

func Start(
	lifecycle fx.Lifecycle,
	setup brokerconfigurate.Brokers, //nolint:all
	interactor interactors.Relay,
) {
	setup.Setup()

	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				interactor.SendMessagesToBroker()
				return nil
			},
			OnStop: func(context.Context) error {
				fmt.Println("Stopping application")
				return nil
			},
		},
	)
}
