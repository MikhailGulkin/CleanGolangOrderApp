package command

import (
	addressRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/address/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/repo"
	productRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/persistence/repo"
	userRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/user/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/address/vo"
	order "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/order/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/order/services"
	vo2 "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/order/vo"
	vo3 "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/product/vo"
	vo4 "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/user/vo"
	"reflect"
)

type CreateOrderImpl struct {
	command.CreateOrder
	productRepo.ProductRepo
	userRepo.UserRepo
	addressRepo.AddressRepo
	repo.OrderRepo
	services.Service
	persistence.UoW
}

func (interactor *CreateOrderImpl) Create(command command.CreateOrderCommand) error {
	products, productError := interactor.ProductRepo.AcquireProductsByIDs(vo3.GetProductIDs(command.ProductsIDs))
	if productError != nil {
		return productError
	}
	user, userError := interactor.UserRepo.AcquireUserByID(vo4.UserID{Value: command.UserID})
	if userError != nil {
		return userError
	}
	address, addressError := interactor.AddressRepo.AcquireAddressByID(vo.AddressID{Value: command.DeliveryAddress})
	if addressError != nil {
		return addressError
	}
	previousOrder, previousOrderError := interactor.OrderRepo.AcquireLastOrder()
	if previousOrderError != nil {
		return previousOrderError
	}
	serialNumber := 0
	if !reflect.ValueOf(previousOrder).IsZero() {
		serialNumber = previousOrder.GetSerialNumber()
	}
	orderAddress, orderErrAddress := order.OrderAddress{}.Create(address.AddressID.Value, address.GetFullAddress())

	if orderErrAddress != nil {
		return orderErrAddress
	}
	client, clientError := order.OrderClient{}.Create(user.UserID.Value, user.Username)
	if clientError != nil {
		return clientError
	}
	orderAggregate, err := interactor.Service.CreateOrder(
		vo2.OrderID{Value: command.OrderID},
		orderAddress,
		client,
		serialNumber,
		products,
	)
	if err != nil {
		return err
	}
	interactor.UoW.StartTx()
	err = interactor.OrderRepo.AddOrder(orderAggregate, interactor.UoW.GetTx())
	if err != nil {
		interactor.UoW.Rollback()
		return err
	}
	err = interactor.UoW.Commit()
	if err != nil {
		return err
	}
	return nil
}
