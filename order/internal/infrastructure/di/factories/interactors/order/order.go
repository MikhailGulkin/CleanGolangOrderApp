package order

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/interfaces/persistence"
	outboxRepo "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/interfaces/persistence/repo"
	ch "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/cache"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/command"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/cache"
	c "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/command"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/persistence/dao"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/persistence/reader"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/persistence/repo"
	q "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/query"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/saga"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/query"
	s "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/saga"
	d "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/relay/interfaces/persistence/dao"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/services"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/logger"
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
func NewGetAllOrders(reader reader.OrderCacheReader) q.GetAllOrders {
	return &query.GetAllOrderImpl{
		OrderCacheReader: reader,
	}
}
func NewGetOrderByID(reader reader.OrderCacheReader) q.GetOrderByID {
	return &query.GetOrderByIDImpl{
		OrderCacheReader: reader,
	}
}

func NewSagaCreateOrder(dao dao.OrderSagaDAO, uow persistence.UoW, logger logger.Logger, outboxDAO d.OutboxDAO) saga.CreateOrder {
	return &s.CreateOrderImpl{
		OrderSagaDAO: dao,
		UoW:          uow,
		Logger:       logger,
		OutboxDAO:    outboxDAO,
	}
}
func NewCacheCreateOrder(cacheDAO dao.OrderCacheDAO, logger logger.Logger) cache.OrderCache {
	return &ch.OrderCacheImpl{OrderCacheDAO: cacheDAO, Logger: logger}
}
func NewGetAllOrdersByUserID(reader reader.OrderCacheReader, logger logger.Logger) q.GetAllOrdersByUserID {
	return &query.GetAllOrdersByUserIDImpl{
		OrderCacheReader: reader,
		Logger:           logger,
	}
}
func NewDeleteOrderByID(
	logger logger.Logger,
	repo repo.OrderRepo,
	uow persistence.UoW,
	outRepo outboxRepo.OutboxRepo,
) c.DeleteOrder {
	return &command.DeleteOrderImpl{
		Logger:     logger,
		OrderRepo:  repo,
		UoW:        uow,
		OutboxRepo: outRepo,
	}
}

var Module = fx.Provide(
	NewCreateOrder,
	NewDeleteOrderByID,

	NewGetAllOrders,
	NewGetOrderByID,

	NewGetAllOrdersByUserID,

	NewCacheCreateOrder,

	NewSagaCreateOrder,
)
