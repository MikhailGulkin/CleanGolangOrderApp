package controllers

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/providers/controllers/handlers"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/providers/controllers/routes"
	"go.uber.org/fx"
)

var Module = fx.Options(
	handlers.Module,
	routes.Module,
)
