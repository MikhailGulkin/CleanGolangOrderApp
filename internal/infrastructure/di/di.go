package di

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/di/factories/cache"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/di/factories/db"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/di/factories/interactors"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/di/factories/logger"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/di/factories/mediator"
	messagebroker "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/di/factories/messageBroker"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"infrastructure.di",
	fx.Options(
		db.Module,
		interactors.Module,
		messagebroker.Module,
		cache.Module,
		logger.Module,
		mediator.Module,
	),
)
