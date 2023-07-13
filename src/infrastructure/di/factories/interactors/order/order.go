package order

import (
	addressRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/address/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/command"
	c "github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/repo"
	productRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/persistence/repo"
	userRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/user/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/services/order"
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
		Service:     order.Service{},
	}
}

var Module = fx.Provide(
	NewCreateOrder,
)
