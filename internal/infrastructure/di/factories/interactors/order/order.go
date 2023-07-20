package order

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/common/interfaces/persistence"
	outboxRepo "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/common/interfaces/persistence/repo"
	ch "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/cache"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/cache"
	c "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/persistence/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/persistence/reader"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/persistence/repo"
	q "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/query"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/saga"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/query"
	s "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/saga"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/order/services"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/logger"
	"go.uber.org/fx"
)

func NewCreateOrder(
	uow persistence.UoW,
	orderRepo repo.OrderRepo,
	orderDAO dao.OrderDAO,
	outboxRepo outboxRepo.OutboxRepo,
) c.CreateOrder {
	return &command.CreateOrderImpl{
		UoW:        uow,
		OrderRepo:  orderRepo,
		Service:    services.Service{},
		OutboxRepo: outboxRepo,
		OrderDAO:   orderDAO,
	}
}
func NewGetAllOrders(dao reader.OrderReader) q.GetAllOrders {
	return &query.GetAllOrderImpl{
		OrderReader: dao,
	}
}
func NewGetOrderByID(dao reader.OrderReader) q.GetOrderByID {
	return &query.GetOrderByIDImpl{
		OrderReader: dao,
	}
}

func NewSagaCreateOrder(dao dao.OrderSagaDAO, uow persistence.UoW, logger logger.Logger) saga.CreateOrder {
	return &s.CreateOrderImpl{
		OrderSagaDAO: dao,
		UoW:          uow,
		Logger:       logger,
	}
}
func NewCacheCreateOrder(cacheDAO dao.OrderCacheDAO) cache.OrderCache {
	return &ch.OrderCacheImpl{OrderCacheDAO: cacheDAO}
}

var Module = fx.Provide(
	NewCreateOrder,
	NewGetAllOrders,
	NewGetOrderByID,

	NewCacheCreateOrder,

	NewSagaCreateOrder,
)
