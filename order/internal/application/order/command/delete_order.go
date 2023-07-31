package command

import (
	"fmt"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/interfaces/logger"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/interfaces/persistence"
	outboxRepo "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/interfaces/persistence/repo"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/command"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/persistence/dao"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/persistence/repo"
)

type DeleteOrderImpl struct {
	command.DeleteOrder
	repo.OrderRepo
	persistence.UoW
	dao.OrderDAO
	outboxRepo.OutboxRepo
	logger.Logger
}

func (interactor *DeleteOrderImpl) Delete(c command.DeleteOrderCommand) error {
	order, err := interactor.OrderRepo.AcquiredOrder(c.OrderID)
	if err != nil {
		return err
	}
	err = order.DeleteOrder()
	if err != nil {
		return err
	}
	err = interactor.OrderDAO.DeleteOrder(c.OrderID, interactor.UoW.StartTx())
	if err != nil {
		interactor.UoW.Rollback()
		return err
	}
	err = interactor.AddEvents(order.PullEvents(), interactor.UoW.GetTx())
	if err != nil {
		return err
	}
	err = interactor.UoW.Commit()
	if err != nil {
		return err
	}
	interactor.Logger.Info(fmt.Sprintf("Delete Event, id %s", c.OrderID.String()))
	return nil
}
