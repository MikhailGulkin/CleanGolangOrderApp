package product

import (
	appDAO "github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/persistence/reader"
	appRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/reader"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/reader/product"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/repo"
	productRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/repo/product"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func BuildProductRepo(conn *gorm.DB) appRepo.ProductRepo {
	return &productRepo.RepoImpl{BaseGormRepo: repo.BaseGormRepo{Session: conn}}
}
func BuildProductReader(conn *gorm.DB) appDAO.ProductReader {
	return &product.ReaderImpl{BaseGormDAO: reader.BaseGormDAO{Session: conn}}
}

var Module = fx.Provide(
	BuildProductRepo,
	BuildProductReader,
)
