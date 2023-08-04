package command

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/interfaces/persistence"
	outbox "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/interfaces/persistence/repo"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/command"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/persistence/dao"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/persistence/repo"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/services"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/vo"
	"reflect"
)

type CreateOrderImpl struct {
	command.CreateOrder
	repo.OrderRepo
	services.Service
	persistence.UoW
	outbox.OutboxRepo
	dao.OrderDAO
}

func (interactor *CreateOrderImpl) Create(command command.CreateOrderCommand) error {
	previousOrder, previousOrderError := interactor.OrderRepo.AcquireLastOrder()
	if previousOrderError != nil {
		return previousOrderError
	}
	serialNumber := 0
	if !reflect.ValueOf(previousOrder).IsZero() {
		serialNumber = previousOrder.SerialNumber
	}

	products, productError := interactor.OrderDAO.GetProductsByIDs(command.ProductsIDs)
	if productError != nil {
		return productError
	}
	orderAggregate, err := interactor.Service.CreateOrder(
		vo.OrderID{Value: command.OrderID},
		command.DeliveryAddress,
		command.UserID,
		serialNumber,
		products,
	)
	if err != nil {
		return err
	}
	// TODO: Придумать что-то с uow и его получением -- нужен новый экземпляр на каждый запрос

	uow := interactor.UoW.Get()

	err = interactor.OrderRepo.AddOrder(&orderAggregate, uow.StartTx())
	if err != nil {
		uow.Rollback()
		return err
	}
	err = interactor.OutboxRepo.AddEvents(orderAggregate.PullEvents(), uow.GetTx())
	if err != nil {
		uow.Rollback()
		return err
	}
	err = uow.Commit()
	if err != nil {
		return err
	}
	return nil
}
