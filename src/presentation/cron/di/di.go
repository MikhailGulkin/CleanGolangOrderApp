package di

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/cron/engine"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/cron/handlers"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/cron/handlers/relay"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	handlers.NewHandlers,
	relay.NewCronHandler,
	engine.NewCron,
	engine.NewCronController,
)
