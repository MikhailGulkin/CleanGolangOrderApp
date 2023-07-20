package saga

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/common/consts"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/common/consts/outbox"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/common/interfaces/logger"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/persistence/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/saga"
	out "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/relay/interfaces/persistence/dao"
)

type CreateOrderImpl struct {
	saga.CreateOrder
	dao.OrderSagaDAO
	out.OutboxDAO
	persistence.UoW
	logger.Logger
}

func (interactor *CreateOrderImpl) CheckStatus(message saga.Message) {
	check, err := interactor.OrderSagaDAO.CheckSagaCompletion(message.OrderID)
	if err != nil || !check {
		return
	}
	switch message.OrderType {
	case consts.Approved:
		err := interactor.OrderSagaDAO.UpdateOrderSagaStatus(message.OrderID, string(consts.Approved), interactor.UoW.StartTx())
		if err != nil {
			interactor.Rollback()
			return
		}
		err = interactor.OutboxDAO.UpdateStatusMessagesByAggregateID(message.OrderID, outbox.Awaiting, interactor.UoW.GetTx())
		if err != nil {
			interactor.Rollback()
			return
		}
		err = interactor.Commit()
		if err != nil {
			return
		}
	case consts.Rejected:
		err := interactor.OrderSagaDAO.DeleteOrderCascade(message.OrderID, string(consts.Rejected), interactor.UoW.StartTx())
		if err != nil {
			interactor.Rollback()
		}
		err = interactor.OutboxDAO.UpdateStatusMessagesByAggregateID(message.OrderID, outbox.Rejected, interactor.UoW.GetTx())
		if err != nil {
			interactor.Rollback()
			return
		}
		err = interactor.UoW.Commit()
		if err != nil {
			interactor.Info("Error when rejected order creation")
			return
		}
	default:
		interactor.Info(fmt.Sprintf("Invalid Order creation saga type %s", message.OrderType))
		return
	}
}
