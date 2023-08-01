package consumer

import (
	"context"
	"fmt"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/logger"
	brokerconfigurate "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/messageBroker/brokerConfigurate"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/consumer/providers"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/consumer/subscribers"
	"go.uber.org/fx"
)

var Module = fx.Options(
	providers.Module,
	fx.Invoke(Start),
)

func Start(
	lifecycle fx.Lifecycle,
	logger logger.Logger,
	consumer subscribers.Subscribers, //nolint:all
	setup brokerconfigurate.Brokers, //nolint:all
) {
	setup.Setup()
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					defer func() {
						if r := recover(); r != nil {
							logger.Info(fmt.Sprintf("Recovered when boot consumer, r %s", r))
						}
						consumer.Listen()
					}()
				}()
				return nil
			},
			OnStop: func(context.Context) error {
				logger.Info("Stopping consumer application")
				return nil
			},
		},
	)
}
