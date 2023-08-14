package aggregate

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/events"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/vo"
	"github.com/google/uuid"
)

func (a *CustomerAggregate) CreateCustomer(
	fullName vo.FullName,
	addressID uuid.UUID,
	balance vo.CustomerBalance,
) error {
	event, err := events.NewCustomerCreatedEvent(a, fullName, addressID, balance)
	if err != nil {
		return err
	}
	return a.Apply(event)
}
func (a *CustomerAggregate) UpdateTransactionCustomer(
	transaction *entities.CustomerTransactions,
) error {
	event, err := events.NewTransactionsUpdatedEvent(a, transaction)
	if err != nil {
		return err
	}
	return a.Apply(event)
}
