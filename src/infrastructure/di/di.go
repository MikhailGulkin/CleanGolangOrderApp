package di

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/db"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/interactors"
	messagebroker "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/message_broker"
	"go.uber.org/fx"
)

var Module = fx.Options(
	db.Module,
	interactors.Module,
	messagebroker.Module,
)
