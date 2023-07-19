package providers

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/consumer/subscribers"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	subscribers.NewEventConsumer,
)