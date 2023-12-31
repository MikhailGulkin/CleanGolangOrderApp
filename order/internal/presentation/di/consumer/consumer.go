package consumer

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/consumer/subscribers"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"presentation.consumer",
	fx.Provide(
		subscribers.NewEventConsumer,
	),
)
