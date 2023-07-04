package db

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/dao"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(BuildConnection),
	fx.Provide(dao.BuildProductDAO),
)
