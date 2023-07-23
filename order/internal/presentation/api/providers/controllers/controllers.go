package controllers

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/providers/controllers/handlers"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/providers/controllers/routes"
	"go.uber.org/fx"
)

var Module = fx.Options(
	handlers.Module,
	routes.Module,
)
