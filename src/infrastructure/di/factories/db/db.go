package db

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/uow"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/db/address"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/di/factories/db/product"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func BuildGormUoW(conn *gorm.DB) persistence.UoW {
	return &uow.GormUoW{Session: conn, Tx: nil}
}

var Module = fx.Options(
	product.Module,
	address.Module,
	fx.Provide(
		BuildGormUoW,
		db.BuildConnection,
	),
)
