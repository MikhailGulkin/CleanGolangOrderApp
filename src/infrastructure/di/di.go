package di

import (
	commandbus "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/commandBus"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/db"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/interactors"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/logger"
	messagebroker "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/messageBroker"
	"go.uber.org/fx"
)

var Module = fx.Options(
	db.Module,
	interactors.Module,
	messagebroker.Module,
	logger.Module,
	commandbus.Module,
)
