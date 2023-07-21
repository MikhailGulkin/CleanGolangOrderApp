package providers

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/consumer/subscribers"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"presentation.consumer",
	fx.Provide(
		subscribers.NewEventConsumer,
	),
)
