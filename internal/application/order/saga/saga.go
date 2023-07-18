package saga

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/common/consts"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/common/interfaces/logger"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/persistence/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/saga"
)

type CreateOrderImpl struct {
	saga.CreateOrder
	dao.OrderSagaDAO
	persistence.UoW
	logger.Logger
}

func (interactor *CreateOrderImpl) CheckStatus(message saga.Message) {
	switch message.OrderType {
	case consts.Approved:
		err := interactor.OrderSagaDAO.UpdateOrderSagaStatus(message.OrderID, string(consts.Approved), interactor.UoW.StartTx())
		if err != nil {
			interactor.Rollback()
		}
		err = interactor.Commit()
		if err != nil {
			return
		}
	case consts.Rejected:
		err := interactor.OrderSagaDAO.DeleteOrderCascade(message.OrderID, interactor.UoW.StartTx())
		if err != nil {
			interactor.Rollback()
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
