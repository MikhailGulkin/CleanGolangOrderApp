package aggregate

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/common"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/consts"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/events"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/vo"
)

func (a *CustomerAggregate) When(evt common.Event) error {

	switch evt.GetEventType() {
	case events.CustomerCreated:
		return a.onCustomerCreated(evt)
	case events.TransactionUpdated:
		return a.onUpdateCustomerTransaction(evt)
	case events.BalanceUpdated:
		return a.onUpdateCustomerBalance(evt)
	default:
		return common.ErrInvalidEventType
	}
}
func (a *CustomerAggregate) onCustomerCreated(evt common.Event) error {
	var eventData events.CustomerCreatedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		return fmt.Errorf("get Json Data, err: %w", err)
	}
	a.Customer.Balance = eventData.Balance
	a.Customer.FullName = eventData.FullName
	a.Customer.AddressID = eventData.AddressID
	a.Customer.Email = eventData.Email
	return nil
}
func (a *CustomerAggregate) onUpdateCustomerTransaction(evt common.Event) error {
	var eventData events.TransactionsUpdatedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		return fmt.Errorf("get Json Data, err: %w", err)
	}
	a.Customer.Transactions = append(a.Customer.Transactions, eventData.Transaction)
	return nil
}
func (a *CustomerAggregate) onUpdateCustomerBalance(evt common.Event) error {
	var eventData events.BalanceUpdatedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		return fmt.Errorf("get Json Data, err: %w", err)
	}
	a.Customer.Balance = eventData.Balance
	return nil
}
func (a *CustomerAggregate) GetNewBalance(money vo.Money, txType consts.TransactionType) vo.Balance {
	balance := a.Customer.Balance
	switch txType {
	case consts.DEPOSIT:
		return balance.DepositBalance(money)
	case consts.PURCHASE_PENDING:
		return balance.Purchase(money)
	default:
		return vo.Balance{}
	}
}
