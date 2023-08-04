package db

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/db"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/db/dao"
	base "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/db/repo"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/db/uow"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/di/factories/db/orders"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/di/factories/db/outbox"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/di/factories/db/product"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func NewBaseRepo(conn *gorm.DB) base.BaseGormRepo {
	return base.BaseGormRepo{Session: conn}
}
func NewBaseDAO(conn *gorm.DB) dao.BaseGormDAO {
	return dao.BaseGormDAO{Session: conn}
}

var Module = fx.Options(
	product.Module,
	orders.Module,
	fx.Provide(
		uow.BuildGormUoW,
		db.BuildConnection,
		NewBaseRepo,
		NewBaseDAO,
	),

	outbox.Module,
)
