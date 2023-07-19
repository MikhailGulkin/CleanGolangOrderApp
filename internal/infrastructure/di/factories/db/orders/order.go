package orders

import (
	appDAO "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/persistence/dao"
	appReader "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/persistence/reader"
	appRepo "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/persistence/repo"
	orderDAO "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/dao/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/reader"
	orderReader "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/reader/order"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/repo/order"
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
	BuildOrderReader,
	BuildOrderDAO,

	BuildOrderSagaDAO,
)