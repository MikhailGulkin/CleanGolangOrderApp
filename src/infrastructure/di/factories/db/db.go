package db

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/uow"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/db/address"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/db/orders"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/db/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/db/user"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func BuildGormUoW(conn *gorm.DB) persistence.UoW {
	return &uow.GormUoW{Session: conn}
}

var Module = fx.Options(
	product.Module,
	address.Module,
	user.Module,
	orders.Module,
	fx.Provide(
		BuildGormUoW,
		db.BuildConnection,
	),
)
