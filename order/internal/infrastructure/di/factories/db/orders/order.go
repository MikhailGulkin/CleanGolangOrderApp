package orders

import (
	appDAO "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/order/interfaces/persistence/dao"
	appRepo "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/order/interfaces/persistence/repo"
	orderDAO "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/dao/order"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/repo/order"
	"go.uber.org/fx"
)

func BuildOrderRepo(base repo.BaseGormRepo) appRepo.OrderRepo {
	return &order.RepoImpl{
		BaseGormRepo: base,
	}
}
func BuildOrderDAO(base repo.BaseGormRepo) appDAO.OrderDAO {
	return &orderDAO.DAOImpl{
		BaseGormRepo: base,
	}
}
func BuildOrderSagaDAO(base repo.BaseGormRepo) appDAO.OrderSagaDAO {
	return &orderDAO.SagaDAOImpl{
		BaseGormRepo: base,
	}
}

var Module = fx.Provide(
	BuildOrderRepo,
	BuildOrderDAO,

	BuildOrderSagaDAO,
)
