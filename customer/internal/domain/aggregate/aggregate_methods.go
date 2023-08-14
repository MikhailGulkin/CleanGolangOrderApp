package aggregate

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/common"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/events"
)

func (a *CustomerAggregate) When(evt common.Event) error {

	switch evt.GetEventType() {
	case events.CustomerCreated:
		return a.onCustomerCreated(evt)
	case events.TransactionUpdated:
		return a.onUpdateCustomerTransaction(evt)
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
