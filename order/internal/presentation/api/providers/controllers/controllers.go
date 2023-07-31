package controllers

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api/providers/controllers/handlers"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api/providers/controllers/routes"
	"go.uber.org/fx"
)

var Module = fx.Options(
	handlers.Module,
	routes.Module,
)
