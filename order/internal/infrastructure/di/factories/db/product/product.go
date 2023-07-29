package product

import (
	appDAO "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/interfaces/persistence/dao"
	appReader "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/interfaces/persistence/reader"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/dao"
	productDAO "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/dao/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/reader/product"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/repo"
	"go.uber.org/fx"
)

func BuildProductRepo(base repo.BaseGormRepo) appDAO.ProductDAO {
	return &productDAO.DAOImpl{BaseGormRepo: base}
}
func BuildProductReader(base dao.BaseGormDAO) appReader.ProductReader {
	return &product.ReaderImpl{BaseGormDAO: base}
}

var Module = fx.Provide(
	BuildProductRepo,
	BuildProductReader,
)
