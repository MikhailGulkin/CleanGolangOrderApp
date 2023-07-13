package providers

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/providers/controllers"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/providers/engine"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/providers/middleware"
	"go.uber.org/fx"
)

var Module = fx.Options(
	middleware.Module,
	engine.Module,
	controllers.Module,
)
