package order

import (
	addressRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/address/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/command"
	c "github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/reader"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/repo"
	q "github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/query"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/query"
	productRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/persistence/repo"
	userRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/user/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/order/services"
	"go.uber.org/fx"
)

func NewCreateOrder(
	userRepo userRepo.UserRepo,
	uow persistence.UoW,
	addressRepo addressRepo.AddressRepo,
	productRepo productRepo.ProductRepo,
	orderRepo repo.OrderRepo,
) c.CreateOrder {
	return &command.CreateOrderImpl{
		UserRepo:    userRepo,
		UoW:         uow,
		AddressRepo: addressRepo,
		ProductRepo: productRepo,
		OrderRepo:   orderRepo,
		Service:     services.Service{},
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

var Module = fx.Provide(
	NewCreateOrder,
	NewGetAllOrders,
	NewGetOrderByID,
)
