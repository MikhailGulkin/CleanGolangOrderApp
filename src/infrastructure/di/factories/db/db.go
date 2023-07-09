package db

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence"
	appDAO "github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/reader"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/reader/product"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/repo"
	productRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/repo/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/uow"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func BuildProductDAO(conn *gorm.DB) appDAO.ProductRepo {
	return &productRepo.RepoImpl{BaseGormRepo: repo.BaseGormRepo{Session: conn}}
}
func BuildProductReader(conn *gorm.DB) appDAO.ProductReader {
	return &product.ReaderImpl{BaseGormDAO: reader.BaseGormDAO{Session: conn}}
}
func BuildGormUoW(conn *gorm.DB) persistence.UoW {
	return &uow.GormUoW{Session: conn, Tx: nil}
}

var Module = fx.Provide(
	BuildProductDAO,
	BuildGormUoW,
	BuildProductReader,
	db.BuildConnection,
)
