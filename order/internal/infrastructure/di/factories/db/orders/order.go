package orders

import (
	appDAO "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/order/interfaces/persistence/dao"
	appRepo "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/order/interfaces/persistence/repo"
	orderDAO "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/dao/order"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/repo/order"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func BuildOrderRepo(conn *gorm.DB) appRepo.OrderRepo {
	return &order.RepoImpl{
		BaseGormRepo: repo.BaseGormRepo{Session: conn},
	}
}
func BuildOrderDAO(conn *gorm.DB) appDAO.OrderDAO {
	return &orderDAO.DAOImpl{
		BaseGormRepo: repo.BaseGormRepo{Session: conn},
	}
}
func BuildOrderSagaDAO(conn *gorm.DB) appDAO.OrderSagaDAO {
	return &orderDAO.SagaDAOImpl{
		BaseGormRepo: repo.BaseGormRepo{Session: conn},
	}
}

var Module = fx.Provide(
	BuildOrderRepo,
	BuildOrderDAO,

	BuildOrderSagaDAO,
)
