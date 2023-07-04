package db

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence"
	appDAO "github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/uow"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func BuildProductDAO(conn *gorm.DB) appDAO.ProductDAO {
	return &dao.ProductDAOImpl{BaseGormDAO: dao.BaseGormDAO{Session: conn}}
}
func BuildGormUoW(conn *gorm.DB) persistence.UoW {
	return &uow.GormUoW{Session: conn, Tx: nil}
}

var Module = fx.Provide(
	BuildProductDAO,
	BuildGormUoW,
	db.BuildConnection,
)