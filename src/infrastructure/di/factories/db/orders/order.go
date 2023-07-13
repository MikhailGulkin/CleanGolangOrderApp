package orders

import (
	appRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/repo"
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

var Module = fx.Provide(
	BuildOrderRepo,
)
