package logger

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/logger"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	logger.NewLogger,
)
