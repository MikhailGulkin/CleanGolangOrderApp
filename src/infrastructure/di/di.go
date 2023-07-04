package di

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/db"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/interactors"
	"go.uber.org/fx"
)

var Module = fx.Options(
	db.Module,
	interactors.Module,
)
