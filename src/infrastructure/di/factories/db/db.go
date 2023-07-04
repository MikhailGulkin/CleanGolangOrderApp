package db

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/dao"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func BuildProductDAO(conn *gorm.DB) persistence.ProductDAO {
	return &dao.ProductDAOImpl{BaseGormDAO: dao.BaseGormDAO{Session: conn}}
}

var Module = fx.Provide(
	BuildProductDAO,
	db.BuildConnection,
)
