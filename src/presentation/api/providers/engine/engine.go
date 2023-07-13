package engine

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/engine"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(engine.NewRequestHandler),
)
