package logger

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/logger"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	logger.NewLogger,
)
