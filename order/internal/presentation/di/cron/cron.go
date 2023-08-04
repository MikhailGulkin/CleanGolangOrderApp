package cron

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/cron/engine"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/cron/handlers"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/cron/handlers/relay"
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
