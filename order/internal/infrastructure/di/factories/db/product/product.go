package product

import (
	appDAO "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/interfaces/persistence/dao"
	appReader "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/interfaces/persistence/reader"
	productDAO "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/dao/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/reader"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/reader/product"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/repo"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func BuildProductRepo(conn *gorm.DB) appDAO.ProductDAO {
	return &productDAO.DAOImpl{BaseGormRepo: repo.BaseGormRepo{Session: conn}}
}
func BuildProductReader(conn *gorm.DB) appReader.ProductReader {
	return &product.ReaderImpl{BaseGormDAO: reader.BaseGormDAO{Session: conn}}
}

var Module = fx.Provide(
	BuildProductRepo,
	BuildProductReader,
)
