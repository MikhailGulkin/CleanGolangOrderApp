package orders

import (
	appReader "github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/reader"
	appRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/reader"
	orderReader "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/reader/order"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/repo/order"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func BuildOrderRepo(conn *gorm.DB) appRepo.OrderRepo {
	return &order.RepoImpl{
		BaseGormRepo: repo.BaseGormRepo{Session: conn},
	}
}
func BuildOrderReader(conn *gorm.DB) appReader.OrderReader {
	return &orderReader.ReaderImpl{
		BaseGormDAO: reader.BaseGormDAO{Session: conn},
	}
}

var Module = fx.Provide(
	BuildOrderRepo,
	BuildOrderReader,
)
