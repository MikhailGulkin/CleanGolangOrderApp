package engine

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/presentation/api/engine"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	engine.NewRequestHandler,
)
