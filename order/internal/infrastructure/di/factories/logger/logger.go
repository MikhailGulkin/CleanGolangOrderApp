package logger

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/logger"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	logger.NewLogger,
)
