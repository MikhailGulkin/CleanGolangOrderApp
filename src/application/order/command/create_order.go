package command

import (
	addressRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/address/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/repo"
	productRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/persistence/repo"
	userRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/user/interfaces/persistence/repo"
	o "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/order"
	service "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/services/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/vo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/vo/product"
	"reflect"
)

type CreateOrderImpl struct { //nolint:maligned
	command.CreateOrder
	productRepo.ProductRepo
	userRepo.UserRepo
	addressRepo.AddressRepo
	repo.OrderRepo
	persistence.UoW
	service service.OrderService
}

func (interactor *CreateOrderImpl) Create(command command.CreateOrderCommand) error {
	products, productError := interactor.ProductRepo.AcquireProductsByIDs(product.GetProductIDs(command.ProductsIDs))
	if productError != nil {
		return productError
	}
	user, userError := interactor.UserRepo.AcquireUserByID(vo.UserID{Value: command.UserID})
	if userError != nil {
		return userError
	}
	address, addressError := interactor.AddressRepo.AcquireAddressByID(vo.AddressID{Value: command.DeliveryAddress})
	if addressError != nil {
		return addressError
	}
	previousOrder, previousOrderError := interactor.OrderRepo.AcquireLastOrderByID(vo.OrderID{Value: command.OrderID})
	if previousOrderError != nil {
		return previousOrderError
	}
	serialNumber := 1
	if !reflect.ValueOf(previousOrder).IsZero() {
		serialNumber = previousOrder.GetSerialNumber()
	}
	orderAddress, orderErrAddress := o.OrderAddress{}.Create(address.BuildingNumber)

	if orderErrAddress != nil {
		return orderErrAddress
	}
	client, clientError := o.OrderClient{}.Create(user.Username)
	if clientError != nil {
		return clientError
	}
	order, err := interactor.service.CreateOrder(
		vo.OrderID{Value: command.OrderID},
		orderAddress,
		client,
		serialNumber,
		products,
	)
	if err != nil {
		return err
	}
	interactor.UoW.StartTx()
	err = interactor.OrderRepo.Add(&order, interactor.UoW.GetTx())
	if err != nil {
		return err
	}
	err = interactor.UoW.Commit()
	if err != nil {
		return err
	}
	return nil
}
