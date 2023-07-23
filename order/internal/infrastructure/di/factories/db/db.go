package db

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/uow"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/di/factories/db/orders"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/di/factories/db/outbox"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/di/factories/db/product"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func BuildGormUoW(conn *gorm.DB) persistence.UoW {
	return &uow.GormUoW{Session: conn}
}

var Module = fx.Options(
	product.Module,
	orders.Module,
	fx.Provide(
		BuildGormUoW,
		db.BuildConnection,
	),

	outbox.Module,
)
