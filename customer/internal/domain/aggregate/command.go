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
	email vo.Email,
) error {
	event, err := events.NewCustomerCreatedEvent(a, fullName, addressID, email, vo.NewBalance())
	if err != nil {
		return err
	}
	return a.Apply(event)
}

func (a *CustomerAggregate) CreateAvatarUri() error {
	event, err := events.NewAvatarUriCreatedEvent(a)
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
	balance := a.GetNewBalance(transaction.TransactionSum, transaction.TransactionType)
	eventBalance, errBalance := events.NewBalanceUpdatedEvent(a, balance)
	if errBalance != nil {
		return errBalance
	}
	if err := a.Apply(eventBalance); err != nil {
		return err
	}
	return a.Apply(event)
}
