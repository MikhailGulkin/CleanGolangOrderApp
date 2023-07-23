package providers

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/presentation/cron/engine"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/presentation/cron/handlers"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/presentation/cron/handlers/relay"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"presentation.cron",
	fx.Provide(
		handlers.NewHandlers,
		relay.NewCronHandler,
		engine.NewCron,
		engine.NewCronController,
	),
)
